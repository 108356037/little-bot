package sniper

import (
	"container/heap"
	"encoding/json"
	"math/big"
	"sync"
	"tigris-bot/price"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/gorilla/websocket"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

func Initialize(record map[int]price.PriceData, c *websocket.Conn, wg *sync.WaitGroup) {
	var received []price.PriceData
	var prices []byte

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Fatal("read error:", err)
		}

		if len(message) > 100 {
			prices = message[10 : len(message)-1]
		}

		err = json.Unmarshal(prices, &received)
		if err != nil {
			log.Printf("Unmarshal error, message: %v", received)
			continue /* Don't do anything if unexpected unmarshal */
		} else {
			for _, priceData := range received {
				if priceData.Signature != "" {
					record[priceData.Asset] = priceData
				}
			}
			wg.Done()
			break
		}
	}
}

func Searcher(record map[int]price.PriceData, c *websocket.Conn, wg *sync.WaitGroup) {
	/* This wg should share the same wg with Initialize */
	wg.Wait()

	var received []price.PriceData
	var prices []byte

	bigIntOne := big.NewInt(1)
	pumpThreshold, _ := decimal.NewFromString("1.001")

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Warn("read error:", err)
			break
		}

		if len(message) > 100 {
			prices = message[10 : len(message)-1]
		} else {
			continue
		}

		err = json.Unmarshal(prices, &received)
		if err != nil {
			log.Printf("Unmarshal error, message: %v", received)
			continue /* Don't do anything if unexpected unmarshal */
		}

		for _, receivedPriceData := range received {
			lastPriceData, ok := record[receivedPriceData.Asset]
			if ok && (receivedPriceData.Signature != "") {
				// lastPrice := new(big.Int)
				// receivedPrice := new(big.Int)
				// pricePumped := new(big.Int)
				// pricePumped := decimal.

				// lastPrice, ok = lastPrice.SetString(lastPriceData.Price, 10)
				// receivedPrice, ok = receivedPrice.SetString(receivedPriceData.Price, 10)

				lastPrice, err := decimal.NewFromString(lastPriceData.Price)
				receivedPrice, err := decimal.NewFromString(receivedPriceData.Price)

				lastTimestamp := new(big.Int)
				receivedTimestamp := new(big.Int)
				timestampDiff := new(big.Int)

				lastTimestamp, ok = lastTimestamp.SetString(lastPriceData.Timestamp, 10)
				receivedTimestamp, ok = receivedTimestamp.SetString(receivedPriceData.Timestamp, 10)

				// if !ok {
				// 	log.Warnf("Error when parsing price, lastPrice: %v, receivedPrice: %v", lastPrice, receivedPrice)
				// 	record[receivedPriceData.Asset] = receivedPriceData

				if err != nil {
					log.Warnf("Error when parsing price, lastPrice: %v, receivedPrice: %v", lastPrice, receivedPrice)
					record[receivedPriceData.Asset] = receivedPriceData
				} else {
					if (timestampDiff.Sub(receivedTimestamp, lastTimestamp)).Cmp(bigIntOne) < 1 {

						pumped := receivedPrice.Div(lastPrice)

						if pumped.Cmp(pumpThreshold) > -1 {
							log.Infof(
								"received: %s, last: %s, asset: %d, receivedTimestamp: %d, lastTimestamp: %d, pumped: %s",
								receivedPrice.String(),
								lastPrice.String(),
								receivedPriceData.Asset,
								receivedTimestamp,
								lastTimestamp,
								pumped.String())
							// go analyzer.LogFollowingPriceData(c)
						}

					}
					record[receivedPriceData.Asset] = receivedPriceData
				}
			} else if receivedPriceData.Signature != "" {
				record[receivedPriceData.Asset] = receivedPriceData
			}
		}
	}
}

func InitializeWithHeapRecord(record map[int]([]price.PriceData), heapifyRecord map[int](*price.PriceDataPriorityQueue), lookbackPriceAmount int, c *websocket.Conn, wg *sync.WaitGroup) {
	var received []price.PriceData
	var prices []byte

	completedAssets := mapset.NewSet[int]()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Fatal("read error:", err)
		}

		if len(message) > 100 {
			prices = message[10 : len(message)-1]
		}

		err = json.Unmarshal(prices, &received)
		if err != nil {
			log.Printf("Unmarshal error, message: %v", received)
			continue /* Don't do anything if unexpected unmarshal */
		} else {
			for _, priceData := range received {
				if (priceData.Signature != "") &&
					(len(record[priceData.Asset]) < lookbackPriceAmount) {
					record[priceData.Asset] = append(record[priceData.Asset], priceData)
				} else if len(record[priceData.Asset]) == lookbackPriceAmount {
					completedAssets.Add(priceData.Asset)
				}
			}

			/* 33 is the number of all assets minus unsuported assets::42-9=33 */
			if len(completedAssets.ToSlice()) == 33 {
				for k, v := range record {
					h := &price.PriceDataPriorityQueue{}
					for _, _v := range v {
						// *h = append(*h, price.ConvertPriceDataForCalc(_v))
						// heap.Push(h, price.ConvertPriceDataForCalc(_v))
						h.Push(price.ConvertToPriceDataForCalc(_v))
					}
					heap.Init(h)
					heapifyRecord[k] = h
				}
				wg.Done()
				log.Info("Done initializing")
				break
			}
		}
	}
}

func SearcherWithHeapRecord(heapifyRecord map[int](*price.PriceDataPriorityQueue), pumpThreshold string, c *websocket.Conn, wg *sync.WaitGroup) {
	/* This wg should share the same wg with Initialize */
	wg.Wait()

	var received []price.PriceData
	var prices []byte

	// bigDecimalOne := decimal.NewFromInt(1)
	bigDecimalTwo := decimal.NewFromInt(2)
	threshold, _ := decimal.NewFromString(pumpThreshold)

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Warn("read error:", err)
			break
		}

		if len(message) > 100 {
			prices = message[10 : len(message)-1]
		} else {
			continue
		}

		err = json.Unmarshal(prices, &received)
		if err != nil {
			log.Printf("Unmarshal error, message: %v", received)
			continue /* Don't do anything if unexpected unmarshal */
		}

		for _, receivedPriceData := range received {
			lastPriceHeap, ok := heapifyRecord[receivedPriceData.Asset]
			if ok && (receivedPriceData.Signature != "") {
				maxAttempts := 3
				receivedPrice, err := decimal.NewFromString(receivedPriceData.Price)
				receivedTimestamp, err := decimal.NewFromString(receivedPriceData.Timestamp)
				if err != nil {
					log.Fatal(err)
				}

				for i := 0; i < maxAttempts; i++ {
					exploitablePrice := heap.Pop(lastPriceHeap).(price.PriceDataForCalc)

					/* Price too old, skip */
					if (receivedTimestamp.Sub(exploitablePrice.Timestamp).Cmp(bigDecimalTwo)) > 0 {
						continue
					}

					pumped := receivedPrice.Div(exploitablePrice.Price)
					if pumped.Cmp(threshold) > -1 {
						log.Infof(
							"received: %s, last: %s, asset: %d, receivedTimestamp: %s, lastTimestamp: %s, pumped: %s",
							receivedPrice.String(),
							exploitablePrice.Price.String(),
							receivedPriceData.Asset,
							receivedTimestamp.String(),
							exploitablePrice.Timestamp.String(),
							pumped.String())
						log.Infof("Current heap length: %d", len(*lastPriceHeap))
					}
					break
				}
				heap.Push(lastPriceHeap, price.ConvertToPriceDataForCalc(receivedPriceData))
			}
		}
	}
}
