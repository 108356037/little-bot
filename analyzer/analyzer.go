package analyzer

import (
	"context"
	"fmt"
	"os"
	"tigris-bot/static"
	"time"

	"github.com/gorilla/websocket"
	"github.com/valyala/fasthttp"
)

var (
	AnalyzerConn *websocket.Conn
	client       *fasthttp.Client
	// mu           map[int]*sync.Mutex
)

func InitAnalyzer() {
	readTimeout, _ := time.ParseDuration("500ms")
	writeTimeout, _ := time.ParseDuration("500ms")
	maxIdleConnDuration, _ := time.ParseDuration("1h")

	client = &fasthttp.Client{
		ReadTimeout:                   readTimeout,
		WriteTimeout:                  writeTimeout,
		MaxIdleConnDuration:           maxIdleConnDuration,
		NoDefaultUserAgentHeader:      true, // Don't send: User-Agent: fasthttp
		DisableHeaderNamesNormalizing: true, // If you set the case on your headers correctly you can enable this
		DisablePathNormalizing:        true,
		// increase DNS cache time to an hour instead of default minute
		Dial: (&fasthttp.TCPDialer{
			Concurrency:      4096,
			DNSCacheDuration: time.Hour,
		}).Dial,
	}

	// mu = new(sync.Mutex)
}

/*
func LogFollowingPriceData(c *websocket.Conn) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for {
		log.Info("we are inside anaylzer")
		select {
		case <-ctx.Done():
			log.Info("Times up, break out of LogFollowingPriceData")
			break
		default:
			go func() {
				for {
					_, message, err := c.ReadMessage()
					if err != nil {
						log.Warn("read error:", err)
						break
					}
					log.Info(message)
				}
			}()
		}
	}
}
*/

func LogFollowingPriceData(traceDuration time.Duration, assetId int) {
	ctx, cancel := context.WithTimeout(context.Background(), traceDuration)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			// log.Info("times up, break out")
			return
		default:
			getPythnetPrice(assetId)
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func getPythnetPrice(assetId int) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("https://xc-mainnet.pyth.network/api/latest_price_feeds")

	req.URI().SetQueryString(`ids[]=` + static.TigrisIdToAssetInfo[assetId].PythnetId)
	// req.URI().QueryArgs().Add(`ids[]`, static.TigrisIDToAssetInfo[assetId].PythnetId)
	// fmt.Println(req.URI().String())

	req.Header.SetMethod(fasthttp.MethodGet)

	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)

	fasthttp.ReleaseRequest(req)
	if err == nil {
		fmt.Printf("DEBUG Response: %s\n", resp.Body())
	} else {
		fmt.Fprintf(os.Stderr, "ERR Connection error: %v\n", err)
	}
	fasthttp.ReleaseResponse(resp)
}
