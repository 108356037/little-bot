package price

import (
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

var (
	DIVISION_CONSTANT decimal.Decimal
)

type RecievedData struct {
	data       string
	priceDatas []PriceData
}

type PriceData struct {
	Signature string `json:"signature"`
	Provider  string `json:"provider"`
	Asset     int    `json:"asset"`
	Timestamp string `json:"timestamp"`
	IsClosed  bool   `json:"is_closed"`
	Price     string `json:"price"`
	Spread    int64  `json:"spread"`
}

type PriceDataForCalc struct {
	Signature string          `json:"signature"`
	Provider  string          `json:"provider"`
	Asset     int             `json:"asset"`
	Timestamp decimal.Decimal `json:"timestamp"`
	IsClosed  bool            `json:"is_closed"`
	Price     decimal.Decimal `json:"price"`
	Spread    int64           `json:"spread"`
}

func InitPrice() {
	_DIVISION_CONSTANT, err := decimal.NewFromString("10000000000") // 1e10
	if err != nil {
		log.Fatal(err)
	}
	DIVISION_CONSTANT = _DIVISION_CONSTANT
}

func ConvertToPriceDataForCalc(pricedata PriceData) PriceDataForCalc {
	t, err := decimal.NewFromString(pricedata.Timestamp)
	p, err := decimal.NewFromString(pricedata.Price)
	if err != nil {
		log.Fatalf("Inside ConvertPriceDataForCalc: %s", err.Error())
	}

	data := PriceDataForCalc{
		Signature: pricedata.Signature,
		Provider:  pricedata.Provider,
		Asset:     pricedata.Asset,
		IsClosed:  pricedata.IsClosed,
		Spread:    pricedata.Spread,
		Timestamp: t,
		Price:     p,
	}

	return data
}

func ConvertToPriceData(priceDataForCalc PriceDataForCalc) PriceData {
	return PriceData{
		Signature: priceDataForCalc.Signature,
		Provider:  priceDataForCalc.Provider,
		Asset:     priceDataForCalc.Asset,
		IsClosed:  priceDataForCalc.IsClosed,
		Spread:    priceDataForCalc.Spread,
		Timestamp: priceDataForCalc.Timestamp.String(),
		Price:     priceDataForCalc.Price.String(),
	}
}

func PriceAfterSpread(price decimal.Decimal, spread decimal.Decimal, position bool) decimal.Decimal {
	if position == true {
		delta := (price.Mul(spread)).Div(DIVISION_CONSTANT)
		return price.Add(delta)
	} else {
		delta := (price.Mul(spread)).Div(DIVISION_CONSTANT)
		return price.Sub(delta)
	}
}

type PriceDataPriorityQueue []PriceDataForCalc

func (pq PriceDataPriorityQueue) Len() int { return len(pq) }

func (pq PriceDataPriorityQueue) Less(i, j int) bool {
	cmp := pq[i].Price.Cmp(pq[j].Price)
	return cmp == -1
}

func (pq PriceDataPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push and Pop use pointer receivers because they modify the slice's length,
// not just its contents.
func (pq *PriceDataPriorityQueue) Push(x any) {
	*pq = append(*pq, x.(PriceDataForCalc))
}

func (pq *PriceDataPriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
