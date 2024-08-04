package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"tigris-bot/analyzer"
	"tigris-bot/order"
	"tigris-bot/price"
	"tigris-bot/sniper"
	"tigris-bot/static"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var (
	/* possible options: beta, eu1, eu2, us1, us2 */
	oracleWss = "wss://eu1.tigrisoracle.net/socket.io/?EIO=4&transport=websocket"

	sniperWg sync.WaitGroup

	lookbackPriceAmount = 13 /*最多留幾筆過去的 priceData*/
)

func main() {
	log.Info(oracleWss)

	price.InitPrice()
	order.InitOrder()
	static.InitStatic()
	analyzer.InitAnalyzer()
	sniper.InitSniper()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	c, _, err := websocket.DefaultDialer.Dial(oracleWss, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	go sniper.WebSocketKeepAlive(time.Second*5, c)

	sniperWg.Add(1)

	/*
		unsupportedAssetId := mapset.NewSet[int]()
		unsupportedAssetId.Add(2)
		unsupportedAssetId.Add(7)
		unsupportedAssetId.Add(8)
		unsupportedAssetId.Add(9)
		unsupportedAssetId.Add(10)
		unsupportedAssetId.Add(12)
		unsupportedAssetId.Add(16)
		unsupportedAssetId.Add(17)
		unsupportedAssetId.Add(22)
		unsupportedAssetId.Add(25)
		unsupportedAssetId.Add(28)
		unsupportedAssetId.Add(31)
		unsupportedAssetId.Add(32)

		// for _, _id := range unsupportedAssetId.ToSlice() {
		// 	log.Infof("unsupported: %s", static.TigrisIdToAssetInfo[_id].Name)
		// }

		fifoRecord := map[int]([]price.PriceDataForCalc){}
		for i := 0; i < 41; i++ {
			if !unsupportedAssetId.Contains(i) {
			fifoRecord[i] = make([]price.PriceDataForCalc, 0, lookbackPriceAmount)
			}
		}
		go sniper.InitializeWithUnsupportedAssetsFifoRecord(fifoRecord, lookbackPriceAmount, unsupportedAssetId, c, &sniperWg)
	*/

	allAssets := mapset.NewSet[int]()
	for i := 0; i < 41; i++ {
		allAssets.Add(i)
	}

	allAssets.Remove(25)
	allAssets.Remove(9)
	allAssets.Remove(31)
	allAssets.Remove(17)
	allAssets.Remove(16)
	allAssets.Remove(22)
	allAssets.Remove(8)
	allAssets.Remove(11)
	allAssets.Remove(28)
	allAssets.Remove(12)
	allAssets.Remove(36) // PEPE
	allAssets.Remove(11) // ETH/BTC
	allAssets.Remove(33) // LINK/BTC
	allAssets.Remove(34) // XMR/BTC
	// targetAssets := mapset.NewSet[int]()
	// targetAssets.Add(0)
	// targetAssets.Add(1)
	// targetAssets.Add(3)
	// targetAssets.Add(4)
	// targetAssets.Add(5)
	// targetAssets.Add(6)
	// targetAssets.Add(13)
	// targetAssets.Add(14)
	// targetAssets.Add(15)
	// targetAssets.Add(18)
	// targetAssets.Add(19)
	// targetAssets.Add(36)

	fifoRecord := map[int]([]price.PriceDataForCalc){}
	for i := 0; i < 41; i++ { // 目前支援到 priceId 41
		if allAssets.Contains(i) {
			fifoRecord[i] = make([]price.PriceDataForCalc, 0, lookbackPriceAmount)
		}
	}
	go sniper.InitializeWithFifoRecord(fifoRecord, lookbackPriceAmount, allAssets, c, &sniperWg)

	if os.Getenv("ENTRY_THRESHOLD") == "" {
		log.Fatal("Env ENTRY_THRESHOLD must be set")
	}
	log.Infof("Entry position threshold: %s", os.Getenv("ENTRY_THRESHOLD"))

	if os.Getenv("POSITION") == "LONG" {
		go sniper.SearcherWithFifoRecordLong(fifoRecord, os.Getenv("ENTRY_THRESHOLD"), c, &sniperWg)
	}

	if os.Getenv("POSITION") == "SHORT" {
		go sniper.SearcherWithFifoRecordShort(fifoRecord, os.Getenv("ENTRY_THRESHOLD"), c, &sniperWg)
	}

	if os.Getenv("POSITION") == "" {
		log.Fatal("Must set env POSITION")
	}

	/*
		heapifyRecord := map[int](*price.PriceDataPriorityQueue){}
		for i := 0; i < 41; i++ {
			if !unsupportedAssetId.Contains(i) {
				heapifyRecord[i] = nil
			}
		}

		go sniper.InitializeWithHeapRecord(fifoRecord, heapifyRecord, lookbackPriceAmount, c, &sniperWg)
		go sniper.SearcherWithHeapRecord(heapifyRecord, "1.0005", c, &sniperWg)
	*/

	/* msg=40 fires the price quote stream */
	err = c.WriteMessage(1, []byte(strconv.Itoa(40)))
	if err != nil {
		log.Fatal("dial:", err)
	}

	go order.GasUpdater(time.Second * 5)
	go order.NonceUpdater(time.Second * 3)

	/*
		cAnalyze, _, err := websocket.DefaultDialer.Dial(oracleWss, nil)
		if err != nil {
			log.Fatal("dial:", err)
		}
		analyzer.AnalyzerConn = cAnalyze
		defer analyzer.AnalyzerConn.Close()
		go sniper.WebSocketKeepAlive(time.Second*5, analyzer.AnalyzerConn)

		err = cAnalyze.WriteMessage(1, []byte(strconv.Itoa(40)))
		if err != nil {
			log.Fatal("dial:", err)
		}
	*/

	/* block until receives os signal */
	<-ctx.Done()
	stop()
}
