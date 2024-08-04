package order

import (
	"math/big"
	"strings"
	"tigris-bot/trading"

	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"

	"tigris-bot/price"

	"github.com/ethereum/go-ethereum/common"
)

func SendOrder(priceData price.PriceData, priceAfterSpread decimal.Decimal, postion bool) {
	provider := common.HexToAddress(priceData.Provider)
	isClosed := priceData.IsClosed
	asset := big.NewInt(int64(priceData.Asset))
	spread := big.NewInt(priceData.Spread)

	price := new(big.Int)
	price, _ = price.SetString(priceData.Price, 10)

	timestamp := new(big.Int)
	timestamp, _ = timestamp.SetString(priceData.Timestamp, 10)

	priceDataToSend := trading.PriceData{
		Provider:  provider,
		IsClosed:  isClosed,
		Asset:     asset,
		Spread:    spread,
		Price:     price,
		Timestamp: timestamp,
	}

	permitDataToSend := trading.ITradingERC20PermitData{
		Deadline:  common.Big0,
		Amount:    common.Big0,
		V:         0,
		R:         [32]byte{},
		S:         [32]byte{},
		UsePermit: false,
	}

	var slPrice decimal.Decimal
	var tpPrice decimal.Decimal
	if postion == true {
		slPrice = priceAfterSpread.Mul(stopLossRatioForLong)
		tpPrice = priceAfterSpread.Mul(takeProfitRatioForLong)
	} else {
		slPrice = priceAfterSpread.Mul(stopLossRatioForShort)
		tpPrice = priceAfterSpread.Mul(takeProfitRatioForShort)
	}

	tradeInfoToSend := trading.ITradingTradeInfo{
		Margin:      fiveUMargin,
		MarginAsset: common.HexToAddress("Fd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9"), // USDT
		StableVault: common.HexToAddress("e82fcefbDD034500B5862B4827CAE5c117f6b921"), // tigrisUSD
		Leverage:    hundredLeverage,
		Asset:       big.NewInt(int64(priceData.Asset)),
		Direction:   postion, // true for long, false for short
		TpPrice:     tpPrice.BigInt(),
		SlPrice:     slPrice.BigInt(),
		Referrer:    common.HexToAddress("604fd04bEfb586F67C8A509eEBd0e5e53DFadA0F"),
	}

	sig := common.Hex2Bytes(strings.TrimPrefix(priceData.Signature, "0x"))
	trader := auth.From
	// trader := common.HexToAddress("bcaBDc5425ef27a83e185857E908b16B79B1F514") //22kventures

	txResult, err := tradingContract.InitiateMarketOrder(auth, tradeInfoToSend, priceDataToSend, sig, permitDataToSend, trader)
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(txResult.Nonce() + 1))
	log.Info(txResult)
}
