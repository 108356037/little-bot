package order

import (
	"context"
	"math/big"
	"time"

	log "github.com/sirupsen/logrus"
)

func GasUpdater(intervalTime time.Duration) {
	ticker := time.NewTicker(intervalTime)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			gas, err := traderClient.SuggestGasPrice(context.Background())
			if err != nil {
				log.Warnf("Cannot get current gas price: %s", err.Error())
			}

			updateGas := new(big.Int)
			scale := big.NewInt(2)

			updateGas.Mul(gas, scale)

			if updateGas.Cmp(gas) != 1 {
				log.Warn("Cannot get scaled gas price")
			} else {
				auth.GasPrice = updateGas
			}
		}
	}
}

func NonceUpdater(intervalTime time.Duration) {
	ticker := time.NewTicker(intervalTime)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			nextNonce, err := traderClient.PendingNonceAt(context.Background(), auth.From)
			if err != nil {
				log.Warnf("Cannot get next nonce: %s", err.Error())
			}
			auth.Nonce = big.NewInt(int64(nextNonce))
		}
	}
}
