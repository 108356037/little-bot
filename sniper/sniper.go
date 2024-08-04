package sniper

import (
	"encoding/json"
	"strconv"
	"sync"
	"time"

	"tigris-bot/order"
	"tigris-bot/price"

	"github.com/shopspring/decimal"

	log "github.com/sirupsen/logrus"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/gorilla/websocket"
)

var (
	mu map[int]*sync.Mutex
)

func InitSniper() {
	_mu := map[int]*sync.Mutex{}
	for i := 0; i < 41; i++ {
		_mu[i] = new(sync.Mutex)
	}
	mu = _mu
}

func WebSocketKeepAlive(intervalTime time.Duration, c *websocket.Conn) {
	ticker := time.NewTicker(intervalTime)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := c.WriteMessage(1, []byte(strconv.Itoa(3)))
			if err != nil {
				log.Warnf("Error when sending message 3 to websocket: %s", err.Error())
			}
			err = c.WriteMessage(websocket.PingMessage, []byte("ping"))
			if err != nil {
				log.Warnf("Error when pinging websocket connection: %s", err.Error())
			}
		}
	}
}

func InitializeWithUnsupportedAssetsFifoRecord(record map[int]([]price.PriceDataForCalc), lookbackPriceAmount int, unsupportedAssets mapset.Set[int], c *websocket.Conn, wg *sync.WaitGroup) {

	isValidAsset := func(priceData price.PriceData) bool {
		return (priceData.Signature != "") &&
			(!unsupportedAssets.Contains(priceData.Asset)) &&
			(len(record[priceData.Asset]) < lookbackPriceAmount)
	}

	totalAssets := 42
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
				if isValidAsset(priceData) {
					record[priceData.Asset] = append(
						record[priceData.Asset], price.ConvertToPriceDataForCalc(priceData))
				} else if len(record[priceData.Asset]) == lookbackPriceAmount {
					completedAssets.Add(priceData.Asset)
					// log.Infof("Completed asset: %d", priceData.Asset)
				}
			}

			// log.Info(completedAssets.Difference(allAssets.Difference(unsupportedAssets)))
			if len(completedAssets.ToSlice()) == totalAssets-len(unsupportedAssets.ToSlice()) {
				wg.Done()
				log.Info("Done initializing")
				break
			}
		}
	}
}

func InitializeWithFifoRecord(record map[int]([]price.PriceDataForCalc), lookbackPriceAmount int, targetAssets mapset.Set[int], c *websocket.Conn, wg *sync.WaitGroup) {
	isValidAsset := func(priceData price.PriceData) bool {
		return (priceData.Signature != "") &&
			(targetAssets.Contains(priceData.Asset)) &&
			(len(record[priceData.Asset]) < lookbackPriceAmount)
	}

	allAssets := mapset.NewSet[int]()
	for i := 0; i < 41; i++ {
		allAssets.Add(i)
	}

	completedAssets := mapset.NewSet[int]()
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
				if isValidAsset(priceData) {
					record[priceData.Asset] = append(
						record[priceData.Asset], price.ConvertToPriceDataForCalc(priceData))
				} else if len(record[priceData.Asset]) == lookbackPriceAmount {
					completedAssets.Add(priceData.Asset)
					// log.Infof("Completed asset: %d", priceData.Asset)
				}
			}

			// log.Info(allAssets.Difference(completedAssets))
			// log.Info(len(completedAssets.ToSlice()) == len(targetAssets.ToSlice()))
			if len(completedAssets.ToSlice()) == len(targetAssets.ToSlice()) {
				wg.Done()
				log.Info("Done initializing")
				break
			}
		}
	}
}

func SearcherWithFifoRecordLong(record map[int]([]price.PriceDataForCalc), pumpThreshold string, c *websocket.Conn, wg *sync.WaitGroup) {
	log.Info("This is long postion searcher")
	/* This wg should share the same wg with Initialize */
	wg.Wait()

	position := true
	var received []price.PriceData
	var prices []byte

	// bigDecimalOne := decimal.NewFromInt(1)
	bigDecimalTwo := decimal.NewFromInt(2)
	// bigDecimalThree := decimal.NewFromInt(3)
	threshold, err := decimal.NewFromString(pumpThreshold)
	if err != nil {
		log.Fatalf("Error on setting pumpThreshold: %s", err.Error())
	}

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Fatalf("Error when reading message from websocket: %s", err.Error())
			break
		}

		if len(message) > 100 {
			prices = message[10 : len(message)-1]
		} else {
			continue /* Price message abnormal, skip this epoch */
		}

		err = json.Unmarshal(prices, &received)
		if err != nil {
			log.Printf("Unmarshal error, message: %v", received)
			continue /* Don't do anything if unexpected unmarshal */
		}

		for _, receivedPriceData := range received {
			lastPriceDatas, ok := record[receivedPriceData.Asset]

			if ok && (receivedPriceData.Signature != "") && (!receivedPriceData.IsClosed) {
				if !mutexLocked(mu[receivedPriceData.Asset]) {
					receivedPrice, err := decimal.NewFromString(receivedPriceData.Price)
					receivedTimestamp, err := decimal.NewFromString(receivedPriceData.Timestamp)
					if err != nil {
						log.Fatal(err)
					}

					for _, lastPriceData := range lastPriceDatas {
						entryPrice := price.PriceAfterSpread(
							lastPriceData.Price, decimal.NewFromInt(lastPriceData.Spread), position)

						timestampDelta := receivedTimestamp.Sub(lastPriceData.Timestamp)
						priceDiffPercent := receivedPrice.Div(entryPrice)
						// priceDiffPercentNoSpread := lastPriceData.Price.Div(receivedPrice)
						// log.Infof("%s", priceDiffPercent.String())
						// log.Infof("no spread: %s", priceDiffPercentNoSpread.String())

						if (timestampDelta.Cmp(bigDecimalTwo) < 1) && (priceDiffPercent.Cmp(threshold) > -1) {
							log.Debugf(
								"received: %s, entryPrice: %s, asset: %d, receivedTimestamp: %s, entryPriceTimestamp: %s, pumped: %s, spread: %d",
								receivedPrice.String(), entryPrice.String(),
								receivedPriceData.Asset,
								receivedTimestamp.String(), lastPriceData.Timestamp.String(),
								priceDiffPercent.String(), receivedPriceData.Spread)

							go lockAndAction(
								price.ConvertToPriceData(lastPriceData),
								entryPrice,
								position)
							break /* found exploitable price, break out */
						}
					}
				}
				record[receivedPriceData.Asset] = updatePriceSlice(
					lastPriceDatas, price.ConvertToPriceDataForCalc(receivedPriceData))
				log.Debugf("Asset Id: %d, length of updatePrices: %d",
					receivedPriceData.Asset, len(record[receivedPriceData.Asset]))
			}
		}
	}
}

func SearcherWithFifoRecordShort(record map[int]([]price.PriceDataForCalc), pumpThreshold string, c *websocket.Conn, wg *sync.WaitGroup) {
	log.Info("This is short postion searcher")
	/* This wg should share the same wg with Initialize */
	wg.Wait()

	var received []price.PriceData
	var prices []byte
	position := false

	// bigDecimalOne := decimal.NewFromInt(1)
	bigDecimalTwo := decimal.NewFromInt(2)
	// bigDecimalThree := decimal.NewFromInt(3)
	threshold, err := decimal.NewFromString(pumpThreshold)
	if err != nil {
		log.Fatalf("Error on setting pumpThreshold: %s", err.Error())
	}

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Fatalf("Error when reading message from websocket: %s", err.Error())
			break
		}

		if len(message) > 100 {
			prices = message[10 : len(message)-1]
		} else {
			continue /* Price message abnormal, skip this epoch */
		}

		err = json.Unmarshal(prices, &received)
		if err != nil {
			log.Printf("Unmarshal error, message: %v", received)
			continue /* Don't do anything if unexpected unmarshal */
		}

		for _, receivedPriceData := range received {
			lastPriceDatas, ok := record[receivedPriceData.Asset]

			if ok && (receivedPriceData.Signature != "") && (!receivedPriceData.IsClosed) {
				if !mutexLocked(mu[receivedPriceData.Asset]) {
					receivedPrice, err := decimal.NewFromString(receivedPriceData.Price)
					receivedTimestamp, err := decimal.NewFromString(receivedPriceData.Timestamp)
					if err != nil {
						log.Fatal(err)
					}

					for _, lastPriceData := range lastPriceDatas {
						entryPrice := price.PriceAfterSpread(
							lastPriceData.Price, decimal.NewFromInt(lastPriceData.Spread), position)
						// log.Infof("entryPrice: %s", entryPrice.String())
						// log.Infof("lastPriceData.Price: %s", lastPriceData.Price.String())

						timestampDelta := receivedTimestamp.Sub(lastPriceData.Timestamp)
						priceDiffPercent := entryPrice.Div(receivedPrice)
						// priceDiffPercentNoSpread := lastPriceData.Price.Div(receivedPrice)
						// log.Infof("%s", priceDiffPercent.String())
						// log.Infof("no spread: %s", priceDiffPercentNoSpread.String())

						if (timestampDelta.Cmp(bigDecimalTwo) < 1) && (priceDiffPercent.Cmp(threshold) > -1) {
							log.Debugf(
								"received: %s, entryPrice: %s, asset: %d, receivedTimestamp: %s, entryPriceTimestamp: %s, pumped: %s, spread: %d",
								receivedPrice.String(), entryPrice.String(),
								receivedPriceData.Asset,
								receivedTimestamp.String(), lastPriceData.Timestamp.String(),
								priceDiffPercent.String(), receivedPriceData.Spread)

							go lockAndAction(
								price.ConvertToPriceData(lastPriceData),
								entryPrice,
								position)
							break /* found exploitable price, break out */
						}
					}
				}
				record[receivedPriceData.Asset] = updatePriceSlice(
					lastPriceDatas, price.ConvertToPriceDataForCalc(receivedPriceData))
				log.Debugf("Asset Id: %d, length of updatePrices: %d",
					receivedPriceData.Asset, len(record[receivedPriceData.Asset]))
				log.Debugf("%v %v", lastPriceDatas[0].Timestamp.String(), receivedPriceData.Timestamp)
			}
		}
	}
}

func lockAndAction(exploitPriceData price.PriceData, exploitPriceAfterSpread decimal.Decimal, position bool) {
	if mu[exploitPriceData.Asset].TryLock() {
		defer mu[exploitPriceData.Asset].Unlock()
		// analyzer.LogFollowingPriceData(time.Second*20, exploitablePriceData.Asset)
		order.SendOrder(exploitPriceData, exploitPriceAfterSpread, position)
		time.Sleep(time.Second * 60)
	}
	return
}
