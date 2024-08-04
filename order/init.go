package order

import (
	"context"
	"crypto/ecdsa"
	"os"

	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"

	"math/big"
	"tigris-bot/trading"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	fiveUString           = "5000000000000000000"
	hundredLeverageString = "100000000000000000000"
)

var (
	traderClient    *ethclient.Client
	tradingContract *trading.Trading

	fiveUMargin     *big.Int
	hundredLeverage *big.Int

	stopLossRatioForLong   decimal.Decimal
	takeProfitRatioForLong decimal.Decimal

	stopLossRatioForShort   decimal.Decimal
	takeProfitRatioForShort decimal.Decimal

	nextNonce *big.Int
	auth      *bind.TransactOpts
)

func InitOrder() {
	gethUrl := os.Getenv("ARB_URL")
	log.Infof("Arb url: %s", gethUrl)
	client, err := ethclient.Dial(gethUrl)
	if err != nil {
		log.Fatal(err)
	}
	traderClient = client

	pk := os.Getenv("PK")
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	nextNonce = big.NewInt(int64(nonce))

	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	_auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(42161))
	if err != nil {
		log.Fatal(err)
	}
	_auth.Nonce = nextNonce
	_auth.Value = common.Big0        // in wei
	_auth.GasLimit = uint64(4500000) // in units
	_auth.GasPrice = big.NewInt(int64(200000000))
	// _auth.GasTipCap = big.NewInt(int64(200000000)) // 2gwei
	// _auth.NoSend = true

	auth = _auth

	log.Printf("%#v", auth)

	instance, err := trading.NewTrading(common.HexToAddress("2B6026d7b69f0fa4e703D965Bb0FEF0Fa838fEad"), traderClient)
	if err != nil {
		log.Fatal(err)
	}
	tradingContract = instance

	_stopLossRatioL, err := decimal.NewFromString("0.9997")
	if err != nil {
		log.Fatal(err)
	}
	stopLossRatioForLong = _stopLossRatioL

	_takeProfitRatioL, err := decimal.NewFromString("1.0035")
	if err != nil {
		log.Fatal(err)
	}
	takeProfitRatioForLong = _takeProfitRatioL

	_stopLossRatioS, err := decimal.NewFromString("1.0003")
	if err != nil {
		log.Fatal(err)
	}
	stopLossRatioForShort = _stopLossRatioS

	_takeProfitRatioS, err := decimal.NewFromString("0.9965")
	if err != nil {
		log.Fatal(err)
	}
	takeProfitRatioForShort = _takeProfitRatioS

	_fiveUMargin := new(big.Int)
	_fiveUMargin, _ = _fiveUMargin.SetString(fiveUString, 10) // 10x
	if _fiveUMargin == nil {
		log.Fatal("Cannot set _fiveUMargin big int")
	}
	fiveUMargin = _fiveUMargin

	_100xLeverage := new(big.Int)
	_100xLeverage, _ = _100xLeverage.SetString(hundredLeverageString, 10) // 10x
	if _100xLeverage == nil {
		log.Fatal("Cannot set _100xLeverage big int")
	}
	hundredLeverage = _100xLeverage

	log.Infof("TP ratio L: %s", takeProfitRatioForLong.String())
	log.Infof("SL ratio L: %s", stopLossRatioForLong.String())

	log.Infof("TP ratio S: %s", takeProfitRatioForShort.String())
	log.Infof("SL ratio S: %s", stopLossRatioForShort.String())
}
