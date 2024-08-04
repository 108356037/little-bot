// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package trading

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ITradingERC20PermitData is an auto generated low-level Go binding around an user-defined struct.
type ITradingERC20PermitData struct {
	Deadline  *big.Int
	Amount    *big.Int
	V         uint8
	R         [32]byte
	S         [32]byte
	UsePermit bool
}

// ITradingTradeInfo is an auto generated low-level Go binding around an user-defined struct.
type ITradingTradeInfo struct {
	Margin      *big.Int
	MarginAsset common.Address
	StableVault common.Address
	Leverage    *big.Int
	Asset       *big.Int
	Direction   bool
	TpPrice     *big.Int
	SlPrice     *big.Int
	Referrer    common.Address
}

// PriceData is an auto generated low-level Go binding around an user-defined struct.
type PriceData struct {
	Provider  common.Address
	IsClosed  bool
	Asset     *big.Int
	Price     *big.Int
	Spread    *big.Int
	Timestamp *big.Int
}

// TradingMetaData contains all meta data concerning the Trading contract.
var TradingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_position\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_staking\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pairsContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bond\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_xtig\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BadClosePercent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadConstructor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadDeposit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadLeverage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadSetter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadStopLoss\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BelowMinPositionSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CloseToMaxPnL\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LimitNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LimitNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LiqThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedInVault\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedPair\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotLiquidatable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotMargin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotProxy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotVault\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TradingPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueNotEqualToMargin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WaitDelay\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMargin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"addMargin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"AddToPosition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tigAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"daoFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"botFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"FeesDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"LimitCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"asset\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"direction\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"openPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lev\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"LimitOrderExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMargin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLeverage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isMarginAdded\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"MarginModified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"closePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payout\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"PositionClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liqPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"PositionLiquidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"marginAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableVault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"asset\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"direction\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"tpPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structITrading.TradeInfo\",\"name\":\"tradeInfo\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marginAfterFees\",\"type\":\"uint256\"}],\"name\":\"PositionOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isTp\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"UpdateTPSL\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_stableVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_marginAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_addMargin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isClosed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"asset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spread\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPriceData\",\"name\":\"_priceData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"usePermit\",\"type\":\"bool\"}],\"internalType\":\"structITrading.ERC20PermitData\",\"name\":\"_permitData\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"}],\"name\":\"addMargin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isClosed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"asset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spread\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPriceData\",\"name\":\"_priceData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_stableVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_marginAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_addMargin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"usePermit\",\"type\":\"bool\"}],\"internalType\":\"structITrading.ERC20PermitData\",\"name\":\"_permitData\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"}],\"name\":\"addToPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedVault\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"approveProxy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bondDistribution\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"}],\"name\":\"cancelLimitOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closeFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"daoFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burnFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refDiscount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"botFees\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isClosed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"asset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spread\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPriceData\",\"name\":\"_priceData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"executeLimitOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_percent\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isClosed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"asset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spread\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPriceData\",\"name\":\"_priceData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_stableVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_outputToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"}],\"name\":\"initiateCloseOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"marginAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableVault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"asset\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"direction\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"tpPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structITrading.TradeInfo\",\"name\":\"_tradeInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_orderType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"usePermit\",\"type\":\"bool\"}],\"internalType\":\"structITrading.ERC20PermitData\",\"name\":\"_permitData\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"}],\"name\":\"initiateLimitOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"marginAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableVault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"asset\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"direction\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"tpPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structITrading.TradeInfo\",\"name\":\"_tradeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isClosed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"asset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spread\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPriceData\",\"name\":\"_priceData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"usePermit\",\"type\":\"bool\"}],\"internalType\":\"structITrading.ERC20PermitData\",\"name\":\"_permitData\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"}],\"name\":\"initiateMarketOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastLimitUpdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_tp\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isClosed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"asset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spread\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPriceData\",\"name\":\"_priceData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"limitClose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitOrderPriceRange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isClosed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"asset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spread\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPriceData\",\"name\":\"_priceData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"liquidatePosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"marginApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxWinPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"daoFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burnFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refDiscount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"botFees\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"proxyApprovals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_stableVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_removeMargin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isClosed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"asset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spread\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPriceData\",\"name\":\"_priceData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"}],\"name\":\"removeMargin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stableVault\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_bool\",\"type\":\"bool\"}],\"name\":\"setAllowedVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_percent\",\"type\":\"uint256\"}],\"name\":\"setBondDistribution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_open\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_daoFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_burnFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_refDiscount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_botFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_percent\",\"type\":\"uint256\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_range\",\"type\":\"uint256\"}],\"name\":\"setLimitOrderPriceRange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxWinPercent\",\"type\":\"uint256\"}],\"name\":\"setMaxWinPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timeDelay\",\"type\":\"uint256\"}],\"name\":\"setTimeDelay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ext\",\"type\":\"address\"}],\"name\":\"setTradingExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_bool\",\"type\":\"bool\"}],\"name\":\"setTrustedForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"timeDelayPassed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"actionType\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_type\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limitPrice\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isClosed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"asset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spread\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPriceData\",\"name\":\"_priceData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"}],\"name\":\"updateTpSl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultFundingPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TradingABI is the input ABI used to generate the binding from.
// Deprecated: Use TradingMetaData.ABI instead.
var TradingABI = TradingMetaData.ABI

// Trading is an auto generated Go binding around an Ethereum contract.
type Trading struct {
	TradingCaller     // Read-only binding to the contract
	TradingTransactor // Write-only binding to the contract
	TradingFilterer   // Log filterer for contract events
}

// TradingCaller is an auto generated read-only Go binding around an Ethereum contract.
type TradingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TradingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TradingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TradingSession struct {
	Contract     *Trading          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TradingCallerSession struct {
	Contract *TradingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TradingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TradingTransactorSession struct {
	Contract     *TradingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TradingRaw is an auto generated low-level Go binding around an Ethereum contract.
type TradingRaw struct {
	Contract *Trading // Generic contract binding to access the raw methods on
}

// TradingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TradingCallerRaw struct {
	Contract *TradingCaller // Generic read-only contract binding to access the raw methods on
}

// TradingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TradingTransactorRaw struct {
	Contract *TradingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrading creates a new instance of Trading, bound to a specific deployed contract.
func NewTrading(address common.Address, backend bind.ContractBackend) (*Trading, error) {
	contract, err := bindTrading(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trading{TradingCaller: TradingCaller{contract: contract}, TradingTransactor: TradingTransactor{contract: contract}, TradingFilterer: TradingFilterer{contract: contract}}, nil
}

// NewTradingCaller creates a new read-only instance of Trading, bound to a specific deployed contract.
func NewTradingCaller(address common.Address, caller bind.ContractCaller) (*TradingCaller, error) {
	contract, err := bindTrading(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TradingCaller{contract: contract}, nil
}

// NewTradingTransactor creates a new write-only instance of Trading, bound to a specific deployed contract.
func NewTradingTransactor(address common.Address, transactor bind.ContractTransactor) (*TradingTransactor, error) {
	contract, err := bindTrading(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TradingTransactor{contract: contract}, nil
}

// NewTradingFilterer creates a new log filterer instance of Trading, bound to a specific deployed contract.
func NewTradingFilterer(address common.Address, filterer bind.ContractFilterer) (*TradingFilterer, error) {
	contract, err := bindTrading(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TradingFilterer{contract: contract}, nil
}

// bindTrading binds a generic wrapper to an already deployed contract.
func bindTrading(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TradingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trading *TradingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trading.Contract.TradingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trading *TradingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trading.Contract.TradingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trading *TradingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trading.Contract.TradingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trading *TradingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trading.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trading *TradingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trading.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trading *TradingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trading.Contract.contract.Transact(opts, method, params...)
}

// AllowedVault is a free data retrieval call binding the contract method 0xccf845e2.
//
// Solidity: function allowedVault(address ) view returns(bool)
func (_Trading *TradingCaller) AllowedVault(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "allowedVault", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedVault is a free data retrieval call binding the contract method 0xccf845e2.
//
// Solidity: function allowedVault(address ) view returns(bool)
func (_Trading *TradingSession) AllowedVault(arg0 common.Address) (bool, error) {
	return _Trading.Contract.AllowedVault(&_Trading.CallOpts, arg0)
}

// AllowedVault is a free data retrieval call binding the contract method 0xccf845e2.
//
// Solidity: function allowedVault(address ) view returns(bool)
func (_Trading *TradingCallerSession) AllowedVault(arg0 common.Address) (bool, error) {
	return _Trading.Contract.AllowedVault(&_Trading.CallOpts, arg0)
}

// BondDistribution is a free data retrieval call binding the contract method 0xead6d69e.
//
// Solidity: function bondDistribution() view returns(uint256)
func (_Trading *TradingCaller) BondDistribution(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "bondDistribution")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BondDistribution is a free data retrieval call binding the contract method 0xead6d69e.
//
// Solidity: function bondDistribution() view returns(uint256)
func (_Trading *TradingSession) BondDistribution() (*big.Int, error) {
	return _Trading.Contract.BondDistribution(&_Trading.CallOpts)
}

// BondDistribution is a free data retrieval call binding the contract method 0xead6d69e.
//
// Solidity: function bondDistribution() view returns(uint256)
func (_Trading *TradingCallerSession) BondDistribution() (*big.Int, error) {
	return _Trading.Contract.BondDistribution(&_Trading.CallOpts)
}

// CloseFees is a free data retrieval call binding the contract method 0xd2b28b48.
//
// Solidity: function closeFees() view returns(uint256 daoFees, uint256 burnFees, uint256 refDiscount, uint256 botFees)
func (_Trading *TradingCaller) CloseFees(opts *bind.CallOpts) (struct {
	DaoFees     *big.Int
	BurnFees    *big.Int
	RefDiscount *big.Int
	BotFees     *big.Int
}, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "closeFees")

	outstruct := new(struct {
		DaoFees     *big.Int
		BurnFees    *big.Int
		RefDiscount *big.Int
		BotFees     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DaoFees = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BurnFees = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RefDiscount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BotFees = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CloseFees is a free data retrieval call binding the contract method 0xd2b28b48.
//
// Solidity: function closeFees() view returns(uint256 daoFees, uint256 burnFees, uint256 refDiscount, uint256 botFees)
func (_Trading *TradingSession) CloseFees() (struct {
	DaoFees     *big.Int
	BurnFees    *big.Int
	RefDiscount *big.Int
	BotFees     *big.Int
}, error) {
	return _Trading.Contract.CloseFees(&_Trading.CallOpts)
}

// CloseFees is a free data retrieval call binding the contract method 0xd2b28b48.
//
// Solidity: function closeFees() view returns(uint256 daoFees, uint256 burnFees, uint256 refDiscount, uint256 botFees)
func (_Trading *TradingCallerSession) CloseFees() (struct {
	DaoFees     *big.Int
	BurnFees    *big.Int
	RefDiscount *big.Int
	BotFees     *big.Int
}, error) {
	return _Trading.Contract.CloseFees(&_Trading.CallOpts)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address _forwarder) view returns(bool)
func (_Trading *TradingCaller) IsTrustedForwarder(opts *bind.CallOpts, _forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "isTrustedForwarder", _forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address _forwarder) view returns(bool)
func (_Trading *TradingSession) IsTrustedForwarder(_forwarder common.Address) (bool, error) {
	return _Trading.Contract.IsTrustedForwarder(&_Trading.CallOpts, _forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address _forwarder) view returns(bool)
func (_Trading *TradingCallerSession) IsTrustedForwarder(_forwarder common.Address) (bool, error) {
	return _Trading.Contract.IsTrustedForwarder(&_Trading.CallOpts, _forwarder)
}

// LastLimitUpdate is a free data retrieval call binding the contract method 0xf267f8d9.
//
// Solidity: function lastLimitUpdate(uint256 ) view returns(uint256)
func (_Trading *TradingCaller) LastLimitUpdate(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "lastLimitUpdate", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastLimitUpdate is a free data retrieval call binding the contract method 0xf267f8d9.
//
// Solidity: function lastLimitUpdate(uint256 ) view returns(uint256)
func (_Trading *TradingSession) LastLimitUpdate(arg0 *big.Int) (*big.Int, error) {
	return _Trading.Contract.LastLimitUpdate(&_Trading.CallOpts, arg0)
}

// LastLimitUpdate is a free data retrieval call binding the contract method 0xf267f8d9.
//
// Solidity: function lastLimitUpdate(uint256 ) view returns(uint256)
func (_Trading *TradingCallerSession) LastLimitUpdate(arg0 *big.Int) (*big.Int, error) {
	return _Trading.Contract.LastLimitUpdate(&_Trading.CallOpts, arg0)
}

// LimitOrderPriceRange is a free data retrieval call binding the contract method 0x3e2fb3cd.
//
// Solidity: function limitOrderPriceRange() view returns(uint256)
func (_Trading *TradingCaller) LimitOrderPriceRange(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "limitOrderPriceRange")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitOrderPriceRange is a free data retrieval call binding the contract method 0x3e2fb3cd.
//
// Solidity: function limitOrderPriceRange() view returns(uint256)
func (_Trading *TradingSession) LimitOrderPriceRange() (*big.Int, error) {
	return _Trading.Contract.LimitOrderPriceRange(&_Trading.CallOpts)
}

// LimitOrderPriceRange is a free data retrieval call binding the contract method 0x3e2fb3cd.
//
// Solidity: function limitOrderPriceRange() view returns(uint256)
func (_Trading *TradingCallerSession) LimitOrderPriceRange() (*big.Int, error) {
	return _Trading.Contract.LimitOrderPriceRange(&_Trading.CallOpts)
}

// MarginApproved is a free data retrieval call binding the contract method 0x88663de3.
//
// Solidity: function marginApproved(address ) view returns(bool)
func (_Trading *TradingCaller) MarginApproved(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "marginApproved", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MarginApproved is a free data retrieval call binding the contract method 0x88663de3.
//
// Solidity: function marginApproved(address ) view returns(bool)
func (_Trading *TradingSession) MarginApproved(arg0 common.Address) (bool, error) {
	return _Trading.Contract.MarginApproved(&_Trading.CallOpts, arg0)
}

// MarginApproved is a free data retrieval call binding the contract method 0x88663de3.
//
// Solidity: function marginApproved(address ) view returns(bool)
func (_Trading *TradingCallerSession) MarginApproved(arg0 common.Address) (bool, error) {
	return _Trading.Contract.MarginApproved(&_Trading.CallOpts, arg0)
}

// MaxWinPercent is a free data retrieval call binding the contract method 0xa57286d1.
//
// Solidity: function maxWinPercent() view returns(uint256)
func (_Trading *TradingCaller) MaxWinPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "maxWinPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWinPercent is a free data retrieval call binding the contract method 0xa57286d1.
//
// Solidity: function maxWinPercent() view returns(uint256)
func (_Trading *TradingSession) MaxWinPercent() (*big.Int, error) {
	return _Trading.Contract.MaxWinPercent(&_Trading.CallOpts)
}

// MaxWinPercent is a free data retrieval call binding the contract method 0xa57286d1.
//
// Solidity: function maxWinPercent() view returns(uint256)
func (_Trading *TradingCallerSession) MaxWinPercent() (*big.Int, error) {
	return _Trading.Contract.MaxWinPercent(&_Trading.CallOpts)
}

// OpenFees is a free data retrieval call binding the contract method 0x08101d29.
//
// Solidity: function openFees() view returns(uint256 daoFees, uint256 burnFees, uint256 refDiscount, uint256 botFees)
func (_Trading *TradingCaller) OpenFees(opts *bind.CallOpts) (struct {
	DaoFees     *big.Int
	BurnFees    *big.Int
	RefDiscount *big.Int
	BotFees     *big.Int
}, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "openFees")

	outstruct := new(struct {
		DaoFees     *big.Int
		BurnFees    *big.Int
		RefDiscount *big.Int
		BotFees     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DaoFees = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BurnFees = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RefDiscount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BotFees = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OpenFees is a free data retrieval call binding the contract method 0x08101d29.
//
// Solidity: function openFees() view returns(uint256 daoFees, uint256 burnFees, uint256 refDiscount, uint256 botFees)
func (_Trading *TradingSession) OpenFees() (struct {
	DaoFees     *big.Int
	BurnFees    *big.Int
	RefDiscount *big.Int
	BotFees     *big.Int
}, error) {
	return _Trading.Contract.OpenFees(&_Trading.CallOpts)
}

// OpenFees is a free data retrieval call binding the contract method 0x08101d29.
//
// Solidity: function openFees() view returns(uint256 daoFees, uint256 burnFees, uint256 refDiscount, uint256 botFees)
func (_Trading *TradingCallerSession) OpenFees() (struct {
	DaoFees     *big.Int
	BurnFees    *big.Int
	RefDiscount *big.Int
	BotFees     *big.Int
}, error) {
	return _Trading.Contract.OpenFees(&_Trading.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trading *TradingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trading *TradingSession) Owner() (common.Address, error) {
	return _Trading.Contract.Owner(&_Trading.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trading *TradingCallerSession) Owner() (common.Address, error) {
	return _Trading.Contract.Owner(&_Trading.CallOpts)
}

// ProxyApprovals is a free data retrieval call binding the contract method 0x88520fb5.
//
// Solidity: function proxyApprovals(address ) view returns(address)
func (_Trading *TradingCaller) ProxyApprovals(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "proxyApprovals", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyApprovals is a free data retrieval call binding the contract method 0x88520fb5.
//
// Solidity: function proxyApprovals(address ) view returns(address)
func (_Trading *TradingSession) ProxyApprovals(arg0 common.Address) (common.Address, error) {
	return _Trading.Contract.ProxyApprovals(&_Trading.CallOpts, arg0)
}

// ProxyApprovals is a free data retrieval call binding the contract method 0x88520fb5.
//
// Solidity: function proxyApprovals(address ) view returns(address)
func (_Trading *TradingCallerSession) ProxyApprovals(arg0 common.Address) (common.Address, error) {
	return _Trading.Contract.ProxyApprovals(&_Trading.CallOpts, arg0)
}

// TimeDelay is a free data retrieval call binding the contract method 0xc9dec361.
//
// Solidity: function timeDelay() view returns(uint256)
func (_Trading *TradingCaller) TimeDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "timeDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeDelay is a free data retrieval call binding the contract method 0xc9dec361.
//
// Solidity: function timeDelay() view returns(uint256)
func (_Trading *TradingSession) TimeDelay() (*big.Int, error) {
	return _Trading.Contract.TimeDelay(&_Trading.CallOpts)
}

// TimeDelay is a free data retrieval call binding the contract method 0xc9dec361.
//
// Solidity: function timeDelay() view returns(uint256)
func (_Trading *TradingCallerSession) TimeDelay() (*big.Int, error) {
	return _Trading.Contract.TimeDelay(&_Trading.CallOpts)
}

// TimeDelayPassed is a free data retrieval call binding the contract method 0xb4c1f9cc.
//
// Solidity: function timeDelayPassed(uint256 ) view returns(uint256 delay, bool actionType)
func (_Trading *TradingCaller) TimeDelayPassed(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Delay      *big.Int
	ActionType bool
}, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "timeDelayPassed", arg0)

	outstruct := new(struct {
		Delay      *big.Int
		ActionType bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Delay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ActionType = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TimeDelayPassed is a free data retrieval call binding the contract method 0xb4c1f9cc.
//
// Solidity: function timeDelayPassed(uint256 ) view returns(uint256 delay, bool actionType)
func (_Trading *TradingSession) TimeDelayPassed(arg0 *big.Int) (struct {
	Delay      *big.Int
	ActionType bool
}, error) {
	return _Trading.Contract.TimeDelayPassed(&_Trading.CallOpts, arg0)
}

// TimeDelayPassed is a free data retrieval call binding the contract method 0xb4c1f9cc.
//
// Solidity: function timeDelayPassed(uint256 ) view returns(uint256 delay, bool actionType)
func (_Trading *TradingCallerSession) TimeDelayPassed(arg0 *big.Int) (struct {
	Delay      *big.Int
	ActionType bool
}, error) {
	return _Trading.Contract.TimeDelayPassed(&_Trading.CallOpts, arg0)
}

// VaultFundingPercent is a free data retrieval call binding the contract method 0x32e35b76.
//
// Solidity: function vaultFundingPercent() view returns(uint256)
func (_Trading *TradingCaller) VaultFundingPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Trading.contract.Call(opts, &out, "vaultFundingPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaultFundingPercent is a free data retrieval call binding the contract method 0x32e35b76.
//
// Solidity: function vaultFundingPercent() view returns(uint256)
func (_Trading *TradingSession) VaultFundingPercent() (*big.Int, error) {
	return _Trading.Contract.VaultFundingPercent(&_Trading.CallOpts)
}

// VaultFundingPercent is a free data retrieval call binding the contract method 0x32e35b76.
//
// Solidity: function vaultFundingPercent() view returns(uint256)
func (_Trading *TradingCallerSession) VaultFundingPercent() (*big.Int, error) {
	return _Trading.Contract.VaultFundingPercent(&_Trading.CallOpts)
}

// AddMargin is a paid mutator transaction binding the contract method 0x4e0ec49b.
//
// Solidity: function addMargin(uint256 _id, address _stableVault, address _marginAsset, uint256 _addMargin, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature, (uint256,uint256,uint8,bytes32,bytes32,bool) _permitData, address _trader) returns()
func (_Trading *TradingTransactor) AddMargin(opts *bind.TransactOpts, _id *big.Int, _stableVault common.Address, _marginAsset common.Address, _addMargin *big.Int, _priceData PriceData, _signature []byte, _permitData ITradingERC20PermitData, _trader common.Address) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "addMargin", _id, _stableVault, _marginAsset, _addMargin, _priceData, _signature, _permitData, _trader)
}

// AddMargin is a paid mutator transaction binding the contract method 0x4e0ec49b.
//
// Solidity: function addMargin(uint256 _id, address _stableVault, address _marginAsset, uint256 _addMargin, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature, (uint256,uint256,uint8,bytes32,bytes32,bool) _permitData, address _trader) returns()
func (_Trading *TradingSession) AddMargin(_id *big.Int, _stableVault common.Address, _marginAsset common.Address, _addMargin *big.Int, _priceData PriceData, _signature []byte, _permitData ITradingERC20PermitData, _trader common.Address) (*types.Transaction, error) {
	return _Trading.Contract.AddMargin(&_Trading.TransactOpts, _id, _stableVault, _marginAsset, _addMargin, _priceData, _signature, _permitData, _trader)
}

// AddMargin is a paid mutator transaction binding the contract method 0x4e0ec49b.
//
// Solidity: function addMargin(uint256 _id, address _stableVault, address _marginAsset, uint256 _addMargin, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature, (uint256,uint256,uint8,bytes32,bytes32,bool) _permitData, address _trader) returns()
func (_Trading *TradingTransactorSession) AddMargin(_id *big.Int, _stableVault common.Address, _marginAsset common.Address, _addMargin *big.Int, _priceData PriceData, _signature []byte, _permitData ITradingERC20PermitData, _trader common.Address) (*types.Transaction, error) {
	return _Trading.Contract.AddMargin(&_Trading.TransactOpts, _id, _stableVault, _marginAsset, _addMargin, _priceData, _signature, _permitData, _trader)
}

// AddToPosition is a paid mutator transaction binding the contract method 0x0786f2a9.
//
// Solidity: function addToPosition(uint256 _id, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature, address _stableVault, address _marginAsset, uint256 _addMargin, (uint256,uint256,uint8,bytes32,bytes32,bool) _permitData, address _trader) returns()
func (_Trading *TradingTransactor) AddToPosition(opts *bind.TransactOpts, _id *big.Int, _priceData PriceData, _signature []byte, _stableVault common.Address, _marginAsset common.Address, _addMargin *big.Int, _permitData ITradingERC20PermitData, _trader common.Address) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "addToPosition", _id, _priceData, _signature, _stableVault, _marginAsset, _addMargin, _permitData, _trader)
}

// AddToPosition is a paid mutator transaction binding the contract method 0x0786f2a9.
//
// Solidity: function addToPosition(uint256 _id, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature, address _stableVault, address _marginAsset, uint256 _addMargin, (uint256,uint256,uint8,bytes32,bytes32,bool) _permitData, address _trader) returns()
func (_Trading *TradingSession) AddToPosition(_id *big.Int, _priceData PriceData, _signature []byte, _stableVault common.Address, _marginAsset common.Address, _addMargin *big.Int, _permitData ITradingERC20PermitData, _trader common.Address) (*types.Transaction, error) {
	return _Trading.Contract.AddToPosition(&_Trading.TransactOpts, _id, _priceData, _signature, _stableVault, _marginAsset, _addMargin, _permitData, _trader)
}

// AddToPosition is a paid mutator transaction binding the contract method 0x0786f2a9.
//
// Solidity: function addToPosition(uint256 _id, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature, address _stableVault, address _marginAsset, uint256 _addMargin, (uint256,uint256,uint8,bytes32,bytes32,bool) _permitData, address _trader) returns()
func (_Trading *TradingTransactorSession) AddToPosition(_id *big.Int, _priceData PriceData, _signature []byte, _stableVault common.Address, _marginAsset common.Address, _addMargin *big.Int, _permitData ITradingERC20PermitData, _trader common.Address) (*types.Transaction, error) {
	return _Trading.Contract.AddToPosition(&_Trading.TransactOpts, _id, _priceData, _signature, _stableVault, _marginAsset, _addMargin, _permitData, _trader)
}

// ApproveProxy is a paid mutator transaction binding the contract method 0xd67170c5.
//
// Solidity: function approveProxy(address _proxy) payable returns()
func (_Trading *TradingTransactor) ApproveProxy(opts *bind.TransactOpts, _proxy common.Address) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "approveProxy", _proxy)
}

// ApproveProxy is a paid mutator transaction binding the contract method 0xd67170c5.
//
// Solidity: function approveProxy(address _proxy) payable returns()
func (_Trading *TradingSession) ApproveProxy(_proxy common.Address) (*types.Transaction, error) {
	return _Trading.Contract.ApproveProxy(&_Trading.TransactOpts, _proxy)
}

// ApproveProxy is a paid mutator transaction binding the contract method 0xd67170c5.
//
// Solidity: function approveProxy(address _proxy) payable returns()
func (_Trading *TradingTransactorSession) ApproveProxy(_proxy common.Address) (*types.Transaction, error) {
	return _Trading.Contract.ApproveProxy(&_Trading.TransactOpts, _proxy)
}

// CancelLimitOrder is a paid mutator transaction binding the contract method 0x5b14f316.
//
// Solidity: function cancelLimitOrder(uint256 _id, address _trader) returns()
func (_Trading *TradingTransactor) CancelLimitOrder(opts *bind.TransactOpts, _id *big.Int, _trader common.Address) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "cancelLimitOrder", _id, _trader)
}

// CancelLimitOrder is a paid mutator transaction binding the contract method 0x5b14f316.
//
// Solidity: function cancelLimitOrder(uint256 _id, address _trader) returns()
func (_Trading *TradingSession) CancelLimitOrder(_id *big.Int, _trader common.Address) (*types.Transaction, error) {
	return _Trading.Contract.CancelLimitOrder(&_Trading.TransactOpts, _id, _trader)
}

// CancelLimitOrder is a paid mutator transaction binding the contract method 0x5b14f316.
//
// Solidity: function cancelLimitOrder(uint256 _id, address _trader) returns()
func (_Trading *TradingTransactorSession) CancelLimitOrder(_id *big.Int, _trader common.Address) (*types.Transaction, error) {
	return _Trading.Contract.CancelLimitOrder(&_Trading.TransactOpts, _id, _trader)
}

// ExecuteLimitOrder is a paid mutator transaction binding the contract method 0xddf2aa73.
//
// Solidity: function executeLimitOrder(uint256 _id, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature) returns()
func (_Trading *TradingTransactor) ExecuteLimitOrder(opts *bind.TransactOpts, _id *big.Int, _priceData PriceData, _signature []byte) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "executeLimitOrder", _id, _priceData, _signature)
}

// ExecuteLimitOrder is a paid mutator transaction binding the contract method 0xddf2aa73.
//
// Solidity: function executeLimitOrder(uint256 _id, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature) returns()
func (_Trading *TradingSession) ExecuteLimitOrder(_id *big.Int, _priceData PriceData, _signature []byte) (*types.Transaction, error) {
	return _Trading.Contract.ExecuteLimitOrder(&_Trading.TransactOpts, _id, _priceData, _signature)
}

// ExecuteLimitOrder is a paid mutator transaction binding the contract method 0xddf2aa73.
//
// Solidity: function executeLimitOrder(uint256 _id, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature) returns()
func (_Trading *TradingTransactorSession) ExecuteLimitOrder(_id *big.Int, _priceData PriceData, _signature []byte) (*types.Transaction, error) {
	return _Trading.Contract.ExecuteLimitOrder(&_Trading.TransactOpts, _id, _priceData, _signature)
}

// InitiateCloseOrder is a paid mutator transaction binding the contract method 0xfbea508c.
//
// Solidity: function initiateCloseOrder(uint256 _id, uint256 _percent, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature, address _stableVault, address _outputToken, address _trader) returns()
func (_Trading *TradingTransactor) InitiateCloseOrder(opts *bind.TransactOpts, _id *big.Int, _percent *big.Int, _priceData PriceData, _signature []byte, _stableVault common.Address, _outputToken common.Address, _trader common.Address) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "initiateCloseOrder", _id, _percent, _priceData, _signature, _stableVault, _outputToken, _trader)
}

// InitiateCloseOrder is a paid mutator transaction binding the contract method 0xfbea508c.
//
// Solidity: function initiateCloseOrder(uint256 _id, uint256 _percent, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature, address _stableVault, address _outputToken, address _trader) returns()
func (_Trading *TradingSession) InitiateCloseOrder(_id *big.Int, _percent *big.Int, _priceData PriceData, _signature []byte, _stableVault common.Address, _outputToken common.Address, _trader common.Address) (*types.Transaction, error) {
	return _Trading.Contract.InitiateCloseOrder(&_Trading.TransactOpts, _id, _percent, _priceData, _signature, _stableVault, _outputToken, _trader)
}

// InitiateCloseOrder is a paid mutator transaction binding the contract method 0xfbea508c.
//
// Solidity: function initiateCloseOrder(uint256 _id, uint256 _percent, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature, address _stableVault, address _outputToken, address _trader) returns()
func (_Trading *TradingTransactorSession) InitiateCloseOrder(_id *big.Int, _percent *big.Int, _priceData PriceData, _signature []byte, _stableVault common.Address, _outputToken common.Address, _trader common.Address) (*types.Transaction, error) {
	return _Trading.Contract.InitiateCloseOrder(&_Trading.TransactOpts, _id, _percent, _priceData, _signature, _stableVault, _outputToken, _trader)
}

// InitiateLimitOrder is a paid mutator transaction binding the contract method 0x62fd91ee.
//
// Solidity: function initiateLimitOrder((uint256,address,address,uint256,uint256,bool,uint256,uint256,address) _tradeInfo, uint256 _orderType, uint256 _price, (uint256,uint256,uint8,bytes32,bytes32,bool) _permitData, address _trader) returns()
func (_Trading *TradingTransactor) InitiateLimitOrder(opts *bind.TransactOpts, _tradeInfo ITradingTradeInfo, _orderType *big.Int, _price *big.Int, _permitData ITradingERC20PermitData, _trader common.Address) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "initiateLimitOrder", _tradeInfo, _orderType, _price, _permitData, _trader)
}

// InitiateLimitOrder is a paid mutator transaction binding the contract method 0x62fd91ee.
//
// Solidity: function initiateLimitOrder((uint256,address,address,uint256,uint256,bool,uint256,uint256,address) _tradeInfo, uint256 _orderType, uint256 _price, (uint256,uint256,uint8,bytes32,bytes32,bool) _permitData, address _trader) returns()
func (_Trading *TradingSession) InitiateLimitOrder(_tradeInfo ITradingTradeInfo, _orderType *big.Int, _price *big.Int, _permitData ITradingERC20PermitData, _trader common.Address) (*types.Transaction, error) {
	return _Trading.Contract.InitiateLimitOrder(&_Trading.TransactOpts, _tradeInfo, _orderType, _price, _permitData, _trader)
}

// InitiateLimitOrder is a paid mutator transaction binding the contract method 0x62fd91ee.
//
// Solidity: function initiateLimitOrder((uint256,address,address,uint256,uint256,bool,uint256,uint256,address) _tradeInfo, uint256 _orderType, uint256 _price, (uint256,uint256,uint8,bytes32,bytes32,bool) _permitData, address _trader) returns()
func (_Trading *TradingTransactorSession) InitiateLimitOrder(_tradeInfo ITradingTradeInfo, _orderType *big.Int, _price *big.Int, _permitData ITradingERC20PermitData, _trader common.Address) (*types.Transaction, error) {
	return _Trading.Contract.InitiateLimitOrder(&_Trading.TransactOpts, _tradeInfo, _orderType, _price, _permitData, _trader)
}

// InitiateMarketOrder is a paid mutator transaction binding the contract method 0x8ef04a51.
//
// Solidity: function initiateMarketOrder((uint256,address,address,uint256,uint256,bool,uint256,uint256,address) _tradeInfo, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature, (uint256,uint256,uint8,bytes32,bytes32,bool) _permitData, address _trader) returns()
func (_Trading *TradingTransactor) InitiateMarketOrder(opts *bind.TransactOpts, _tradeInfo ITradingTradeInfo, _priceData PriceData, _signature []byte, _permitData ITradingERC20PermitData, _trader common.Address) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "initiateMarketOrder", _tradeInfo, _priceData, _signature, _permitData, _trader)
}

// InitiateMarketOrder is a paid mutator transaction binding the contract method 0x8ef04a51.
//
// Solidity: function initiateMarketOrder((uint256,address,address,uint256,uint256,bool,uint256,uint256,address) _tradeInfo, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature, (uint256,uint256,uint8,bytes32,bytes32,bool) _permitData, address _trader) returns()
func (_Trading *TradingSession) InitiateMarketOrder(_tradeInfo ITradingTradeInfo, _priceData PriceData, _signature []byte, _permitData ITradingERC20PermitData, _trader common.Address) (*types.Transaction, error) {
	return _Trading.Contract.InitiateMarketOrder(&_Trading.TransactOpts, _tradeInfo, _priceData, _signature, _permitData, _trader)
}

// InitiateMarketOrder is a paid mutator transaction binding the contract method 0x8ef04a51.
//
// Solidity: function initiateMarketOrder((uint256,address,address,uint256,uint256,bool,uint256,uint256,address) _tradeInfo, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature, (uint256,uint256,uint8,bytes32,bytes32,bool) _permitData, address _trader) returns()
func (_Trading *TradingTransactorSession) InitiateMarketOrder(_tradeInfo ITradingTradeInfo, _priceData PriceData, _signature []byte, _permitData ITradingERC20PermitData, _trader common.Address) (*types.Transaction, error) {
	return _Trading.Contract.InitiateMarketOrder(&_Trading.TransactOpts, _tradeInfo, _priceData, _signature, _permitData, _trader)
}

// LimitClose is a paid mutator transaction binding the contract method 0x559efc9d.
//
// Solidity: function limitClose(uint256 _id, bool _tp, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature) returns()
func (_Trading *TradingTransactor) LimitClose(opts *bind.TransactOpts, _id *big.Int, _tp bool, _priceData PriceData, _signature []byte) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "limitClose", _id, _tp, _priceData, _signature)
}

// LimitClose is a paid mutator transaction binding the contract method 0x559efc9d.
//
// Solidity: function limitClose(uint256 _id, bool _tp, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature) returns()
func (_Trading *TradingSession) LimitClose(_id *big.Int, _tp bool, _priceData PriceData, _signature []byte) (*types.Transaction, error) {
	return _Trading.Contract.LimitClose(&_Trading.TransactOpts, _id, _tp, _priceData, _signature)
}

// LimitClose is a paid mutator transaction binding the contract method 0x559efc9d.
//
// Solidity: function limitClose(uint256 _id, bool _tp, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature) returns()
func (_Trading *TradingTransactorSession) LimitClose(_id *big.Int, _tp bool, _priceData PriceData, _signature []byte) (*types.Transaction, error) {
	return _Trading.Contract.LimitClose(&_Trading.TransactOpts, _id, _tp, _priceData, _signature)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0x6b6bdd8d.
//
// Solidity: function liquidatePosition(uint256 _id, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature) returns()
func (_Trading *TradingTransactor) LiquidatePosition(opts *bind.TransactOpts, _id *big.Int, _priceData PriceData, _signature []byte) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "liquidatePosition", _id, _priceData, _signature)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0x6b6bdd8d.
//
// Solidity: function liquidatePosition(uint256 _id, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature) returns()
func (_Trading *TradingSession) LiquidatePosition(_id *big.Int, _priceData PriceData, _signature []byte) (*types.Transaction, error) {
	return _Trading.Contract.LiquidatePosition(&_Trading.TransactOpts, _id, _priceData, _signature)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0x6b6bdd8d.
//
// Solidity: function liquidatePosition(uint256 _id, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature) returns()
func (_Trading *TradingTransactorSession) LiquidatePosition(_id *big.Int, _priceData PriceData, _signature []byte) (*types.Transaction, error) {
	return _Trading.Contract.LiquidatePosition(&_Trading.TransactOpts, _id, _priceData, _signature)
}

// RemoveMargin is a paid mutator transaction binding the contract method 0xd1420e60.
//
// Solidity: function removeMargin(uint256 _id, address _stableVault, address _outputToken, uint256 _removeMargin, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature, address _trader) returns()
func (_Trading *TradingTransactor) RemoveMargin(opts *bind.TransactOpts, _id *big.Int, _stableVault common.Address, _outputToken common.Address, _removeMargin *big.Int, _priceData PriceData, _signature []byte, _trader common.Address) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "removeMargin", _id, _stableVault, _outputToken, _removeMargin, _priceData, _signature, _trader)
}

// RemoveMargin is a paid mutator transaction binding the contract method 0xd1420e60.
//
// Solidity: function removeMargin(uint256 _id, address _stableVault, address _outputToken, uint256 _removeMargin, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature, address _trader) returns()
func (_Trading *TradingSession) RemoveMargin(_id *big.Int, _stableVault common.Address, _outputToken common.Address, _removeMargin *big.Int, _priceData PriceData, _signature []byte, _trader common.Address) (*types.Transaction, error) {
	return _Trading.Contract.RemoveMargin(&_Trading.TransactOpts, _id, _stableVault, _outputToken, _removeMargin, _priceData, _signature, _trader)
}

// RemoveMargin is a paid mutator transaction binding the contract method 0xd1420e60.
//
// Solidity: function removeMargin(uint256 _id, address _stableVault, address _outputToken, uint256 _removeMargin, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature, address _trader) returns()
func (_Trading *TradingTransactorSession) RemoveMargin(_id *big.Int, _stableVault common.Address, _outputToken common.Address, _removeMargin *big.Int, _priceData PriceData, _signature []byte, _trader common.Address) (*types.Transaction, error) {
	return _Trading.Contract.RemoveMargin(&_Trading.TransactOpts, _id, _stableVault, _outputToken, _removeMargin, _priceData, _signature, _trader)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Trading *TradingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Trading *TradingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Trading.Contract.RenounceOwnership(&_Trading.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Trading *TradingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Trading.Contract.RenounceOwnership(&_Trading.TransactOpts)
}

// SetAllowedVault is a paid mutator transaction binding the contract method 0x6bbdc24b.
//
// Solidity: function setAllowedVault(address _stableVault, bool _bool) returns()
func (_Trading *TradingTransactor) SetAllowedVault(opts *bind.TransactOpts, _stableVault common.Address, _bool bool) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "setAllowedVault", _stableVault, _bool)
}

// SetAllowedVault is a paid mutator transaction binding the contract method 0x6bbdc24b.
//
// Solidity: function setAllowedVault(address _stableVault, bool _bool) returns()
func (_Trading *TradingSession) SetAllowedVault(_stableVault common.Address, _bool bool) (*types.Transaction, error) {
	return _Trading.Contract.SetAllowedVault(&_Trading.TransactOpts, _stableVault, _bool)
}

// SetAllowedVault is a paid mutator transaction binding the contract method 0x6bbdc24b.
//
// Solidity: function setAllowedVault(address _stableVault, bool _bool) returns()
func (_Trading *TradingTransactorSession) SetAllowedVault(_stableVault common.Address, _bool bool) (*types.Transaction, error) {
	return _Trading.Contract.SetAllowedVault(&_Trading.TransactOpts, _stableVault, _bool)
}

// SetBondDistribution is a paid mutator transaction binding the contract method 0x59f4d828.
//
// Solidity: function setBondDistribution(uint256 _percent) returns()
func (_Trading *TradingTransactor) SetBondDistribution(opts *bind.TransactOpts, _percent *big.Int) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "setBondDistribution", _percent)
}

// SetBondDistribution is a paid mutator transaction binding the contract method 0x59f4d828.
//
// Solidity: function setBondDistribution(uint256 _percent) returns()
func (_Trading *TradingSession) SetBondDistribution(_percent *big.Int) (*types.Transaction, error) {
	return _Trading.Contract.SetBondDistribution(&_Trading.TransactOpts, _percent)
}

// SetBondDistribution is a paid mutator transaction binding the contract method 0x59f4d828.
//
// Solidity: function setBondDistribution(uint256 _percent) returns()
func (_Trading *TradingTransactorSession) SetBondDistribution(_percent *big.Int) (*types.Transaction, error) {
	return _Trading.Contract.SetBondDistribution(&_Trading.TransactOpts, _percent)
}

// SetFees is a paid mutator transaction binding the contract method 0x1d296c16.
//
// Solidity: function setFees(bool _open, uint256 _daoFees, uint256 _burnFees, uint256 _refDiscount, uint256 _botFees, uint256 _percent) returns()
func (_Trading *TradingTransactor) SetFees(opts *bind.TransactOpts, _open bool, _daoFees *big.Int, _burnFees *big.Int, _refDiscount *big.Int, _botFees *big.Int, _percent *big.Int) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "setFees", _open, _daoFees, _burnFees, _refDiscount, _botFees, _percent)
}

// SetFees is a paid mutator transaction binding the contract method 0x1d296c16.
//
// Solidity: function setFees(bool _open, uint256 _daoFees, uint256 _burnFees, uint256 _refDiscount, uint256 _botFees, uint256 _percent) returns()
func (_Trading *TradingSession) SetFees(_open bool, _daoFees *big.Int, _burnFees *big.Int, _refDiscount *big.Int, _botFees *big.Int, _percent *big.Int) (*types.Transaction, error) {
	return _Trading.Contract.SetFees(&_Trading.TransactOpts, _open, _daoFees, _burnFees, _refDiscount, _botFees, _percent)
}

// SetFees is a paid mutator transaction binding the contract method 0x1d296c16.
//
// Solidity: function setFees(bool _open, uint256 _daoFees, uint256 _burnFees, uint256 _refDiscount, uint256 _botFees, uint256 _percent) returns()
func (_Trading *TradingTransactorSession) SetFees(_open bool, _daoFees *big.Int, _burnFees *big.Int, _refDiscount *big.Int, _botFees *big.Int, _percent *big.Int) (*types.Transaction, error) {
	return _Trading.Contract.SetFees(&_Trading.TransactOpts, _open, _daoFees, _burnFees, _refDiscount, _botFees, _percent)
}

// SetLimitOrderPriceRange is a paid mutator transaction binding the contract method 0x686d94da.
//
// Solidity: function setLimitOrderPriceRange(uint256 _range) returns()
func (_Trading *TradingTransactor) SetLimitOrderPriceRange(opts *bind.TransactOpts, _range *big.Int) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "setLimitOrderPriceRange", _range)
}

// SetLimitOrderPriceRange is a paid mutator transaction binding the contract method 0x686d94da.
//
// Solidity: function setLimitOrderPriceRange(uint256 _range) returns()
func (_Trading *TradingSession) SetLimitOrderPriceRange(_range *big.Int) (*types.Transaction, error) {
	return _Trading.Contract.SetLimitOrderPriceRange(&_Trading.TransactOpts, _range)
}

// SetLimitOrderPriceRange is a paid mutator transaction binding the contract method 0x686d94da.
//
// Solidity: function setLimitOrderPriceRange(uint256 _range) returns()
func (_Trading *TradingTransactorSession) SetLimitOrderPriceRange(_range *big.Int) (*types.Transaction, error) {
	return _Trading.Contract.SetLimitOrderPriceRange(&_Trading.TransactOpts, _range)
}

// SetMaxWinPercent is a paid mutator transaction binding the contract method 0x895164e8.
//
// Solidity: function setMaxWinPercent(uint256 _maxWinPercent) returns()
func (_Trading *TradingTransactor) SetMaxWinPercent(opts *bind.TransactOpts, _maxWinPercent *big.Int) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "setMaxWinPercent", _maxWinPercent)
}

// SetMaxWinPercent is a paid mutator transaction binding the contract method 0x895164e8.
//
// Solidity: function setMaxWinPercent(uint256 _maxWinPercent) returns()
func (_Trading *TradingSession) SetMaxWinPercent(_maxWinPercent *big.Int) (*types.Transaction, error) {
	return _Trading.Contract.SetMaxWinPercent(&_Trading.TransactOpts, _maxWinPercent)
}

// SetMaxWinPercent is a paid mutator transaction binding the contract method 0x895164e8.
//
// Solidity: function setMaxWinPercent(uint256 _maxWinPercent) returns()
func (_Trading *TradingTransactorSession) SetMaxWinPercent(_maxWinPercent *big.Int) (*types.Transaction, error) {
	return _Trading.Contract.SetMaxWinPercent(&_Trading.TransactOpts, _maxWinPercent)
}

// SetTimeDelay is a paid mutator transaction binding the contract method 0x39af6ba9.
//
// Solidity: function setTimeDelay(uint256 _timeDelay) payable returns()
func (_Trading *TradingTransactor) SetTimeDelay(opts *bind.TransactOpts, _timeDelay *big.Int) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "setTimeDelay", _timeDelay)
}

// SetTimeDelay is a paid mutator transaction binding the contract method 0x39af6ba9.
//
// Solidity: function setTimeDelay(uint256 _timeDelay) payable returns()
func (_Trading *TradingSession) SetTimeDelay(_timeDelay *big.Int) (*types.Transaction, error) {
	return _Trading.Contract.SetTimeDelay(&_Trading.TransactOpts, _timeDelay)
}

// SetTimeDelay is a paid mutator transaction binding the contract method 0x39af6ba9.
//
// Solidity: function setTimeDelay(uint256 _timeDelay) payable returns()
func (_Trading *TradingTransactorSession) SetTimeDelay(_timeDelay *big.Int) (*types.Transaction, error) {
	return _Trading.Contract.SetTimeDelay(&_Trading.TransactOpts, _timeDelay)
}

// SetTradingExtension is a paid mutator transaction binding the contract method 0x162f8c89.
//
// Solidity: function setTradingExtension(address _ext) returns()
func (_Trading *TradingTransactor) SetTradingExtension(opts *bind.TransactOpts, _ext common.Address) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "setTradingExtension", _ext)
}

// SetTradingExtension is a paid mutator transaction binding the contract method 0x162f8c89.
//
// Solidity: function setTradingExtension(address _ext) returns()
func (_Trading *TradingSession) SetTradingExtension(_ext common.Address) (*types.Transaction, error) {
	return _Trading.Contract.SetTradingExtension(&_Trading.TransactOpts, _ext)
}

// SetTradingExtension is a paid mutator transaction binding the contract method 0x162f8c89.
//
// Solidity: function setTradingExtension(address _ext) returns()
func (_Trading *TradingTransactorSession) SetTradingExtension(_ext common.Address) (*types.Transaction, error) {
	return _Trading.Contract.SetTradingExtension(&_Trading.TransactOpts, _ext)
}

// SetTrustedForwarder is a paid mutator transaction binding the contract method 0xe691d03b.
//
// Solidity: function setTrustedForwarder(address _forwarder, bool _bool) returns()
func (_Trading *TradingTransactor) SetTrustedForwarder(opts *bind.TransactOpts, _forwarder common.Address, _bool bool) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "setTrustedForwarder", _forwarder, _bool)
}

// SetTrustedForwarder is a paid mutator transaction binding the contract method 0xe691d03b.
//
// Solidity: function setTrustedForwarder(address _forwarder, bool _bool) returns()
func (_Trading *TradingSession) SetTrustedForwarder(_forwarder common.Address, _bool bool) (*types.Transaction, error) {
	return _Trading.Contract.SetTrustedForwarder(&_Trading.TransactOpts, _forwarder, _bool)
}

// SetTrustedForwarder is a paid mutator transaction binding the contract method 0xe691d03b.
//
// Solidity: function setTrustedForwarder(address _forwarder, bool _bool) returns()
func (_Trading *TradingTransactorSession) SetTrustedForwarder(_forwarder common.Address, _bool bool) (*types.Transaction, error) {
	return _Trading.Contract.SetTrustedForwarder(&_Trading.TransactOpts, _forwarder, _bool)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Trading *TradingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Trading *TradingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Trading.Contract.TransferOwnership(&_Trading.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Trading *TradingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Trading.Contract.TransferOwnership(&_Trading.TransactOpts, newOwner)
}

// UpdateTpSl is a paid mutator transaction binding the contract method 0xdd75381b.
//
// Solidity: function updateTpSl(bool _type, uint256 _id, uint256 _limitPrice, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature, address _trader) returns()
func (_Trading *TradingTransactor) UpdateTpSl(opts *bind.TransactOpts, _type bool, _id *big.Int, _limitPrice *big.Int, _priceData PriceData, _signature []byte, _trader common.Address) (*types.Transaction, error) {
	return _Trading.contract.Transact(opts, "updateTpSl", _type, _id, _limitPrice, _priceData, _signature, _trader)
}

// UpdateTpSl is a paid mutator transaction binding the contract method 0xdd75381b.
//
// Solidity: function updateTpSl(bool _type, uint256 _id, uint256 _limitPrice, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature, address _trader) returns()
func (_Trading *TradingSession) UpdateTpSl(_type bool, _id *big.Int, _limitPrice *big.Int, _priceData PriceData, _signature []byte, _trader common.Address) (*types.Transaction, error) {
	return _Trading.Contract.UpdateTpSl(&_Trading.TransactOpts, _type, _id, _limitPrice, _priceData, _signature, _trader)
}

// UpdateTpSl is a paid mutator transaction binding the contract method 0xdd75381b.
//
// Solidity: function updateTpSl(bool _type, uint256 _id, uint256 _limitPrice, (address,bool,uint256,uint256,uint256,uint256) _priceData, bytes _signature, address _trader) returns()
func (_Trading *TradingTransactorSession) UpdateTpSl(_type bool, _id *big.Int, _limitPrice *big.Int, _priceData PriceData, _signature []byte, _trader common.Address) (*types.Transaction, error) {
	return _Trading.Contract.UpdateTpSl(&_Trading.TransactOpts, _type, _id, _limitPrice, _priceData, _signature, _trader)
}

// TradingAddToPositionIterator is returned from FilterAddToPosition and is used to iterate over the raw logs and unpacked data for AddToPosition events raised by the Trading contract.
type TradingAddToPositionIterator struct {
	Event *TradingAddToPosition // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TradingAddToPositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingAddToPosition)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TradingAddToPosition)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TradingAddToPositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingAddToPositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingAddToPosition represents a AddToPosition event raised by the Trading contract.
type TradingAddToPosition struct {
	Id        *big.Int
	NewMargin *big.Int
	NewPrice  *big.Int
	AddMargin *big.Int
	Trader    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddToPosition is a free log retrieval operation binding the contract event 0xc1ad8200bf0a20b4a494ddd6d1b21f41731affd12a3c6efcd30a20b2b8643649.
//
// Solidity: event AddToPosition(uint256 id, uint256 newMargin, uint256 newPrice, uint256 addMargin, address trader)
func (_Trading *TradingFilterer) FilterAddToPosition(opts *bind.FilterOpts) (*TradingAddToPositionIterator, error) {

	logs, sub, err := _Trading.contract.FilterLogs(opts, "AddToPosition")
	if err != nil {
		return nil, err
	}
	return &TradingAddToPositionIterator{contract: _Trading.contract, event: "AddToPosition", logs: logs, sub: sub}, nil
}

// WatchAddToPosition is a free log subscription operation binding the contract event 0xc1ad8200bf0a20b4a494ddd6d1b21f41731affd12a3c6efcd30a20b2b8643649.
//
// Solidity: event AddToPosition(uint256 id, uint256 newMargin, uint256 newPrice, uint256 addMargin, address trader)
func (_Trading *TradingFilterer) WatchAddToPosition(opts *bind.WatchOpts, sink chan<- *TradingAddToPosition) (event.Subscription, error) {

	logs, sub, err := _Trading.contract.WatchLogs(opts, "AddToPosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingAddToPosition)
				if err := _Trading.contract.UnpackLog(event, "AddToPosition", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddToPosition is a log parse operation binding the contract event 0xc1ad8200bf0a20b4a494ddd6d1b21f41731affd12a3c6efcd30a20b2b8643649.
//
// Solidity: event AddToPosition(uint256 id, uint256 newMargin, uint256 newPrice, uint256 addMargin, address trader)
func (_Trading *TradingFilterer) ParseAddToPosition(log types.Log) (*TradingAddToPosition, error) {
	event := new(TradingAddToPosition)
	if err := _Trading.contract.UnpackLog(event, "AddToPosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingFeesDistributedIterator is returned from FilterFeesDistributed and is used to iterate over the raw logs and unpacked data for FeesDistributed events raised by the Trading contract.
type TradingFeesDistributedIterator struct {
	Event *TradingFeesDistributed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TradingFeesDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingFeesDistributed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TradingFeesDistributed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TradingFeesDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingFeesDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingFeesDistributed represents a FeesDistributed event raised by the Trading contract.
type TradingFeesDistributed struct {
	TigAsset common.Address
	DaoFees  *big.Int
	BurnFees *big.Int
	RefFees  *big.Int
	BotFees  *big.Int
	Referrer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeesDistributed is a free log retrieval operation binding the contract event 0xb89c76a3dead86dde74a2a7f80bbe6ee35632e71fa4f60fb76177cce95a3cb72.
//
// Solidity: event FeesDistributed(address tigAsset, uint256 daoFees, uint256 burnFees, uint256 refFees, uint256 botFees, address referrer)
func (_Trading *TradingFilterer) FilterFeesDistributed(opts *bind.FilterOpts) (*TradingFeesDistributedIterator, error) {

	logs, sub, err := _Trading.contract.FilterLogs(opts, "FeesDistributed")
	if err != nil {
		return nil, err
	}
	return &TradingFeesDistributedIterator{contract: _Trading.contract, event: "FeesDistributed", logs: logs, sub: sub}, nil
}

// WatchFeesDistributed is a free log subscription operation binding the contract event 0xb89c76a3dead86dde74a2a7f80bbe6ee35632e71fa4f60fb76177cce95a3cb72.
//
// Solidity: event FeesDistributed(address tigAsset, uint256 daoFees, uint256 burnFees, uint256 refFees, uint256 botFees, address referrer)
func (_Trading *TradingFilterer) WatchFeesDistributed(opts *bind.WatchOpts, sink chan<- *TradingFeesDistributed) (event.Subscription, error) {

	logs, sub, err := _Trading.contract.WatchLogs(opts, "FeesDistributed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingFeesDistributed)
				if err := _Trading.contract.UnpackLog(event, "FeesDistributed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeesDistributed is a log parse operation binding the contract event 0xb89c76a3dead86dde74a2a7f80bbe6ee35632e71fa4f60fb76177cce95a3cb72.
//
// Solidity: event FeesDistributed(address tigAsset, uint256 daoFees, uint256 burnFees, uint256 refFees, uint256 botFees, address referrer)
func (_Trading *TradingFilterer) ParseFeesDistributed(log types.Log) (*TradingFeesDistributed, error) {
	event := new(TradingFeesDistributed)
	if err := _Trading.contract.UnpackLog(event, "FeesDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingLimitCancelledIterator is returned from FilterLimitCancelled and is used to iterate over the raw logs and unpacked data for LimitCancelled events raised by the Trading contract.
type TradingLimitCancelledIterator struct {
	Event *TradingLimitCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TradingLimitCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingLimitCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TradingLimitCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TradingLimitCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingLimitCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingLimitCancelled represents a LimitCancelled event raised by the Trading contract.
type TradingLimitCancelled struct {
	Id     *big.Int
	Trader common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLimitCancelled is a free log retrieval operation binding the contract event 0xe90efd55ba607192707390e73dedf6a4d9de8bdc5425d4795458f849b8efe770.
//
// Solidity: event LimitCancelled(uint256 id, address trader)
func (_Trading *TradingFilterer) FilterLimitCancelled(opts *bind.FilterOpts) (*TradingLimitCancelledIterator, error) {

	logs, sub, err := _Trading.contract.FilterLogs(opts, "LimitCancelled")
	if err != nil {
		return nil, err
	}
	return &TradingLimitCancelledIterator{contract: _Trading.contract, event: "LimitCancelled", logs: logs, sub: sub}, nil
}

// WatchLimitCancelled is a free log subscription operation binding the contract event 0xe90efd55ba607192707390e73dedf6a4d9de8bdc5425d4795458f849b8efe770.
//
// Solidity: event LimitCancelled(uint256 id, address trader)
func (_Trading *TradingFilterer) WatchLimitCancelled(opts *bind.WatchOpts, sink chan<- *TradingLimitCancelled) (event.Subscription, error) {

	logs, sub, err := _Trading.contract.WatchLogs(opts, "LimitCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingLimitCancelled)
				if err := _Trading.contract.UnpackLog(event, "LimitCancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLimitCancelled is a log parse operation binding the contract event 0xe90efd55ba607192707390e73dedf6a4d9de8bdc5425d4795458f849b8efe770.
//
// Solidity: event LimitCancelled(uint256 id, address trader)
func (_Trading *TradingFilterer) ParseLimitCancelled(log types.Log) (*TradingLimitCancelled, error) {
	event := new(TradingLimitCancelled)
	if err := _Trading.contract.UnpackLog(event, "LimitCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingLimitOrderExecutedIterator is returned from FilterLimitOrderExecuted and is used to iterate over the raw logs and unpacked data for LimitOrderExecuted events raised by the Trading contract.
type TradingLimitOrderExecutedIterator struct {
	Event *TradingLimitOrderExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TradingLimitOrderExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingLimitOrderExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TradingLimitOrderExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TradingLimitOrderExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingLimitOrderExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingLimitOrderExecuted represents a LimitOrderExecuted event raised by the Trading contract.
type TradingLimitOrderExecuted struct {
	Asset     *big.Int
	Direction bool
	OpenPrice *big.Int
	Lev       *big.Int
	Margin    *big.Int
	Id        *big.Int
	Trader    common.Address
	Executor  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLimitOrderExecuted is a free log retrieval operation binding the contract event 0xfb0417fa63f05f525846fa10af759ba8373825e6789268e86e244d9c71432ddb.
//
// Solidity: event LimitOrderExecuted(uint256 asset, bool direction, uint256 openPrice, uint256 lev, uint256 margin, uint256 id, address trader, address executor)
func (_Trading *TradingFilterer) FilterLimitOrderExecuted(opts *bind.FilterOpts) (*TradingLimitOrderExecutedIterator, error) {

	logs, sub, err := _Trading.contract.FilterLogs(opts, "LimitOrderExecuted")
	if err != nil {
		return nil, err
	}
	return &TradingLimitOrderExecutedIterator{contract: _Trading.contract, event: "LimitOrderExecuted", logs: logs, sub: sub}, nil
}

// WatchLimitOrderExecuted is a free log subscription operation binding the contract event 0xfb0417fa63f05f525846fa10af759ba8373825e6789268e86e244d9c71432ddb.
//
// Solidity: event LimitOrderExecuted(uint256 asset, bool direction, uint256 openPrice, uint256 lev, uint256 margin, uint256 id, address trader, address executor)
func (_Trading *TradingFilterer) WatchLimitOrderExecuted(opts *bind.WatchOpts, sink chan<- *TradingLimitOrderExecuted) (event.Subscription, error) {

	logs, sub, err := _Trading.contract.WatchLogs(opts, "LimitOrderExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingLimitOrderExecuted)
				if err := _Trading.contract.UnpackLog(event, "LimitOrderExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLimitOrderExecuted is a log parse operation binding the contract event 0xfb0417fa63f05f525846fa10af759ba8373825e6789268e86e244d9c71432ddb.
//
// Solidity: event LimitOrderExecuted(uint256 asset, bool direction, uint256 openPrice, uint256 lev, uint256 margin, uint256 id, address trader, address executor)
func (_Trading *TradingFilterer) ParseLimitOrderExecuted(log types.Log) (*TradingLimitOrderExecuted, error) {
	event := new(TradingLimitOrderExecuted)
	if err := _Trading.contract.UnpackLog(event, "LimitOrderExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingMarginModifiedIterator is returned from FilterMarginModified and is used to iterate over the raw logs and unpacked data for MarginModified events raised by the Trading contract.
type TradingMarginModifiedIterator struct {
	Event *TradingMarginModified // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TradingMarginModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingMarginModified)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TradingMarginModified)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TradingMarginModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingMarginModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingMarginModified represents a MarginModified event raised by the Trading contract.
type TradingMarginModified struct {
	Id            *big.Int
	NewMargin     *big.Int
	NewLeverage   *big.Int
	IsMarginAdded bool
	Trader        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMarginModified is a free log retrieval operation binding the contract event 0x2d641b4a49d53c455ccb289f89e2647f2b49b74992092796d42a9a6a60c6dc0f.
//
// Solidity: event MarginModified(uint256 id, uint256 newMargin, uint256 newLeverage, bool isMarginAdded, address trader)
func (_Trading *TradingFilterer) FilterMarginModified(opts *bind.FilterOpts) (*TradingMarginModifiedIterator, error) {

	logs, sub, err := _Trading.contract.FilterLogs(opts, "MarginModified")
	if err != nil {
		return nil, err
	}
	return &TradingMarginModifiedIterator{contract: _Trading.contract, event: "MarginModified", logs: logs, sub: sub}, nil
}

// WatchMarginModified is a free log subscription operation binding the contract event 0x2d641b4a49d53c455ccb289f89e2647f2b49b74992092796d42a9a6a60c6dc0f.
//
// Solidity: event MarginModified(uint256 id, uint256 newMargin, uint256 newLeverage, bool isMarginAdded, address trader)
func (_Trading *TradingFilterer) WatchMarginModified(opts *bind.WatchOpts, sink chan<- *TradingMarginModified) (event.Subscription, error) {

	logs, sub, err := _Trading.contract.WatchLogs(opts, "MarginModified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingMarginModified)
				if err := _Trading.contract.UnpackLog(event, "MarginModified", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarginModified is a log parse operation binding the contract event 0x2d641b4a49d53c455ccb289f89e2647f2b49b74992092796d42a9a6a60c6dc0f.
//
// Solidity: event MarginModified(uint256 id, uint256 newMargin, uint256 newLeverage, bool isMarginAdded, address trader)
func (_Trading *TradingFilterer) ParseMarginModified(log types.Log) (*TradingMarginModified, error) {
	event := new(TradingMarginModified)
	if err := _Trading.contract.UnpackLog(event, "MarginModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Trading contract.
type TradingOwnershipTransferredIterator struct {
	Event *TradingOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TradingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TradingOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TradingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingOwnershipTransferred represents a OwnershipTransferred event raised by the Trading contract.
type TradingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Trading *TradingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TradingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Trading.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TradingOwnershipTransferredIterator{contract: _Trading.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Trading *TradingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TradingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Trading.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingOwnershipTransferred)
				if err := _Trading.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Trading *TradingFilterer) ParseOwnershipTransferred(log types.Log) (*TradingOwnershipTransferred, error) {
	event := new(TradingOwnershipTransferred)
	if err := _Trading.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingPositionClosedIterator is returned from FilterPositionClosed and is used to iterate over the raw logs and unpacked data for PositionClosed events raised by the Trading contract.
type TradingPositionClosedIterator struct {
	Event *TradingPositionClosed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TradingPositionClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingPositionClosed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TradingPositionClosed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TradingPositionClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingPositionClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingPositionClosed represents a PositionClosed event raised by the Trading contract.
type TradingPositionClosed struct {
	Id         *big.Int
	ClosePrice *big.Int
	Percent    *big.Int
	Payout     *big.Int
	Trader     common.Address
	Executor   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPositionClosed is a free log retrieval operation binding the contract event 0x867238cdae52442a9a7a2c9f057ab630df4afa9f3edf9b7c57ea88f4a720482e.
//
// Solidity: event PositionClosed(uint256 id, uint256 closePrice, uint256 percent, uint256 payout, address trader, address executor)
func (_Trading *TradingFilterer) FilterPositionClosed(opts *bind.FilterOpts) (*TradingPositionClosedIterator, error) {

	logs, sub, err := _Trading.contract.FilterLogs(opts, "PositionClosed")
	if err != nil {
		return nil, err
	}
	return &TradingPositionClosedIterator{contract: _Trading.contract, event: "PositionClosed", logs: logs, sub: sub}, nil
}

// WatchPositionClosed is a free log subscription operation binding the contract event 0x867238cdae52442a9a7a2c9f057ab630df4afa9f3edf9b7c57ea88f4a720482e.
//
// Solidity: event PositionClosed(uint256 id, uint256 closePrice, uint256 percent, uint256 payout, address trader, address executor)
func (_Trading *TradingFilterer) WatchPositionClosed(opts *bind.WatchOpts, sink chan<- *TradingPositionClosed) (event.Subscription, error) {

	logs, sub, err := _Trading.contract.WatchLogs(opts, "PositionClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingPositionClosed)
				if err := _Trading.contract.UnpackLog(event, "PositionClosed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePositionClosed is a log parse operation binding the contract event 0x867238cdae52442a9a7a2c9f057ab630df4afa9f3edf9b7c57ea88f4a720482e.
//
// Solidity: event PositionClosed(uint256 id, uint256 closePrice, uint256 percent, uint256 payout, address trader, address executor)
func (_Trading *TradingFilterer) ParsePositionClosed(log types.Log) (*TradingPositionClosed, error) {
	event := new(TradingPositionClosed)
	if err := _Trading.contract.UnpackLog(event, "PositionClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingPositionLiquidatedIterator is returned from FilterPositionLiquidated and is used to iterate over the raw logs and unpacked data for PositionLiquidated events raised by the Trading contract.
type TradingPositionLiquidatedIterator struct {
	Event *TradingPositionLiquidated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TradingPositionLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingPositionLiquidated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TradingPositionLiquidated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TradingPositionLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingPositionLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingPositionLiquidated represents a PositionLiquidated event raised by the Trading contract.
type TradingPositionLiquidated struct {
	Id       *big.Int
	LiqPrice *big.Int
	Trader   common.Address
	Executor common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPositionLiquidated is a free log retrieval operation binding the contract event 0x88bcc1eb0320aca3f26bf2f8c042066d8ee13b4fb6aad33aa8bdb93fa507a2d8.
//
// Solidity: event PositionLiquidated(uint256 id, uint256 liqPrice, address trader, address executor)
func (_Trading *TradingFilterer) FilterPositionLiquidated(opts *bind.FilterOpts) (*TradingPositionLiquidatedIterator, error) {

	logs, sub, err := _Trading.contract.FilterLogs(opts, "PositionLiquidated")
	if err != nil {
		return nil, err
	}
	return &TradingPositionLiquidatedIterator{contract: _Trading.contract, event: "PositionLiquidated", logs: logs, sub: sub}, nil
}

// WatchPositionLiquidated is a free log subscription operation binding the contract event 0x88bcc1eb0320aca3f26bf2f8c042066d8ee13b4fb6aad33aa8bdb93fa507a2d8.
//
// Solidity: event PositionLiquidated(uint256 id, uint256 liqPrice, address trader, address executor)
func (_Trading *TradingFilterer) WatchPositionLiquidated(opts *bind.WatchOpts, sink chan<- *TradingPositionLiquidated) (event.Subscription, error) {

	logs, sub, err := _Trading.contract.WatchLogs(opts, "PositionLiquidated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingPositionLiquidated)
				if err := _Trading.contract.UnpackLog(event, "PositionLiquidated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePositionLiquidated is a log parse operation binding the contract event 0x88bcc1eb0320aca3f26bf2f8c042066d8ee13b4fb6aad33aa8bdb93fa507a2d8.
//
// Solidity: event PositionLiquidated(uint256 id, uint256 liqPrice, address trader, address executor)
func (_Trading *TradingFilterer) ParsePositionLiquidated(log types.Log) (*TradingPositionLiquidated, error) {
	event := new(TradingPositionLiquidated)
	if err := _Trading.contract.UnpackLog(event, "PositionLiquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingPositionOpenedIterator is returned from FilterPositionOpened and is used to iterate over the raw logs and unpacked data for PositionOpened events raised by the Trading contract.
type TradingPositionOpenedIterator struct {
	Event *TradingPositionOpened // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TradingPositionOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingPositionOpened)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TradingPositionOpened)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TradingPositionOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingPositionOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingPositionOpened represents a PositionOpened event raised by the Trading contract.
type TradingPositionOpened struct {
	TradeInfo       ITradingTradeInfo
	OrderType       *big.Int
	Price           *big.Int
	Id              *big.Int
	Trader          common.Address
	MarginAfterFees *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPositionOpened is a free log retrieval operation binding the contract event 0x0b80f30c85c72afa6f48b80c56e5bc11965697da95ff40ad4c3efb734af236f6.
//
// Solidity: event PositionOpened((uint256,address,address,uint256,uint256,bool,uint256,uint256,address) tradeInfo, uint256 orderType, uint256 price, uint256 id, address trader, uint256 marginAfterFees)
func (_Trading *TradingFilterer) FilterPositionOpened(opts *bind.FilterOpts) (*TradingPositionOpenedIterator, error) {

	logs, sub, err := _Trading.contract.FilterLogs(opts, "PositionOpened")
	if err != nil {
		return nil, err
	}
	return &TradingPositionOpenedIterator{contract: _Trading.contract, event: "PositionOpened", logs: logs, sub: sub}, nil
}

// WatchPositionOpened is a free log subscription operation binding the contract event 0x0b80f30c85c72afa6f48b80c56e5bc11965697da95ff40ad4c3efb734af236f6.
//
// Solidity: event PositionOpened((uint256,address,address,uint256,uint256,bool,uint256,uint256,address) tradeInfo, uint256 orderType, uint256 price, uint256 id, address trader, uint256 marginAfterFees)
func (_Trading *TradingFilterer) WatchPositionOpened(opts *bind.WatchOpts, sink chan<- *TradingPositionOpened) (event.Subscription, error) {

	logs, sub, err := _Trading.contract.WatchLogs(opts, "PositionOpened")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingPositionOpened)
				if err := _Trading.contract.UnpackLog(event, "PositionOpened", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePositionOpened is a log parse operation binding the contract event 0x0b80f30c85c72afa6f48b80c56e5bc11965697da95ff40ad4c3efb734af236f6.
//
// Solidity: event PositionOpened((uint256,address,address,uint256,uint256,bool,uint256,uint256,address) tradeInfo, uint256 orderType, uint256 price, uint256 id, address trader, uint256 marginAfterFees)
func (_Trading *TradingFilterer) ParsePositionOpened(log types.Log) (*TradingPositionOpened, error) {
	event := new(TradingPositionOpened)
	if err := _Trading.contract.UnpackLog(event, "PositionOpened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingUpdateTPSLIterator is returned from FilterUpdateTPSL and is used to iterate over the raw logs and unpacked data for UpdateTPSL events raised by the Trading contract.
type TradingUpdateTPSLIterator struct {
	Event *TradingUpdateTPSL // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TradingUpdateTPSLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingUpdateTPSL)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TradingUpdateTPSL)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TradingUpdateTPSLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingUpdateTPSLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingUpdateTPSL represents a UpdateTPSL event raised by the Trading contract.
type TradingUpdateTPSL struct {
	Id     *big.Int
	IsTp   bool
	Price  *big.Int
	Trader common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdateTPSL is a free log retrieval operation binding the contract event 0xe6e02ba81ec882f0c6fb53844fa075c80d4a23c7a9ca64a679664f226493cde7.
//
// Solidity: event UpdateTPSL(uint256 id, bool isTp, uint256 price, address trader)
func (_Trading *TradingFilterer) FilterUpdateTPSL(opts *bind.FilterOpts) (*TradingUpdateTPSLIterator, error) {

	logs, sub, err := _Trading.contract.FilterLogs(opts, "UpdateTPSL")
	if err != nil {
		return nil, err
	}
	return &TradingUpdateTPSLIterator{contract: _Trading.contract, event: "UpdateTPSL", logs: logs, sub: sub}, nil
}

// WatchUpdateTPSL is a free log subscription operation binding the contract event 0xe6e02ba81ec882f0c6fb53844fa075c80d4a23c7a9ca64a679664f226493cde7.
//
// Solidity: event UpdateTPSL(uint256 id, bool isTp, uint256 price, address trader)
func (_Trading *TradingFilterer) WatchUpdateTPSL(opts *bind.WatchOpts, sink chan<- *TradingUpdateTPSL) (event.Subscription, error) {

	logs, sub, err := _Trading.contract.WatchLogs(opts, "UpdateTPSL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingUpdateTPSL)
				if err := _Trading.contract.UnpackLog(event, "UpdateTPSL", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateTPSL is a log parse operation binding the contract event 0xe6e02ba81ec882f0c6fb53844fa075c80d4a23c7a9ca64a679664f226493cde7.
//
// Solidity: event UpdateTPSL(uint256 id, bool isTp, uint256 price, address trader)
func (_Trading *TradingFilterer) ParseUpdateTPSL(log types.Log) (*TradingUpdateTPSL, error) {
	event := new(TradingUpdateTPSL)
	if err := _Trading.contract.UnpackLog(event, "UpdateTPSL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
