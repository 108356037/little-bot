package sniper

import (
	"reflect"
	"sync"
	"tigris-bot/price"

	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

const (
	mutexLockedFlag = 1
)

func getBigDecimalPrices(priceDatas []*price.PriceData) []decimal.Decimal {
	result := make([]decimal.Decimal, 6)

	for _, v := range priceDatas {
		dec, err := decimal.NewFromString(v.Price)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, dec)
	}

	return result
}

func getBigDecimalTimestamp(priceDatas []*price.PriceData) []decimal.Decimal {
	result := make([]decimal.Decimal, 6)

	for _, v := range priceDatas {
		dec, err := decimal.NewFromString(v.Timestamp)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, dec)
	}

	return result
}

func getBigDecimalPricesAndTimestamps(priceDatas []price.PriceData) [][2]decimal.Decimal {
	result := make([][2]decimal.Decimal, len(priceDatas))

	for idx, val := range priceDatas {
		t, err := decimal.NewFromString(val.Timestamp)
		p, err := decimal.NewFromString(val.Price)
		if err != nil {
			log.Fatal(err)
		}
		result[idx][0] = t
		result[idx][1] = p
	}

	return result
}

func updatePriceSlice(priceSlice []price.PriceDataForCalc, newPriceData price.PriceDataForCalc) []price.PriceDataForCalc {
	// keptPriceData := priceSlice[1:]
	keptPriceData := append(priceSlice[1:], newPriceData)
	return keptPriceData
}

func mutexLocked(m *sync.Mutex) bool {
	state := reflect.ValueOf(m).Elem().FieldByName("state")
	return state.Int()&mutexLockedFlag == mutexLockedFlag
}
