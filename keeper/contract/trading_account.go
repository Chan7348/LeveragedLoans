// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
)

// TradingAccountMetaData contains all meta data concerning the TradingAccount contract.
var TradingAccountMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"CancelLimitOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemCTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemUnderlyingAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"closeAll\",\"type\":\"bool\"}],\"name\":\"CloseLong\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemCTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemUnderlyingAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"closeAll\",\"type\":\"bool\"}],\"name\":\"CloseShort\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isETH\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cTokenAmount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dealPrice\",\"type\":\"uint256\"}],\"name\":\"ExecuteLimitOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limitPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expireAt\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"}],\"name\":\"LimitOpenLong\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdcSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limitPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expireAt\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"}],\"name\":\"LimitOpenShort\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldPendingOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"NewPendingOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cTokenMint\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"OpenLong\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdcSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cTokenMint\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"OpenShort\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isETH\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"underlyingAmount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cUSDC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"cancelLimitOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"usdcAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"closeAll\",\"type\":\"bool\"}],\"name\":\"closeLong\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"closeAll\",\"type\":\"bool\"}],\"name\":\"closeShort\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"comptroller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositUSDC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"executeLimitOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getLimitOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCanceled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"openSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expireAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"weth_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usdc_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uniswapV2Pair_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cETH_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cUSDC_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"comptroller_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastCEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastCUsdcBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastOrderId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ethSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expireAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"}],\"name\":\"limitOpenLong\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"usdcSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expireAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"}],\"name\":\"limitOpenShort\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ethSize\",\"type\":\"uint256\"}],\"name\":\"openLong\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"usdcSize\",\"type\":\"uint256\"}],\"name\":\"openShort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeed\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"setPendingOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV2Call\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Pair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cEthAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cUsdcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usdcAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawUSDC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// TradingAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use TradingAccountMetaData.ABI instead.
var TradingAccountABI = TradingAccountMetaData.ABI

// TradingAccount is an auto generated Go binding around an Ethereum contract.
type TradingAccount struct {
	TradingAccountCaller     // Read-only binding to the contract
	TradingAccountTransactor // Write-only binding to the contract
	TradingAccountFilterer   // Log filterer for contract events
}

// TradingAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type TradingAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradingAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TradingAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradingAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TradingAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradingAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TradingAccountSession struct {
	Contract     *TradingAccount   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradingAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TradingAccountCallerSession struct {
	Contract *TradingAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TradingAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TradingAccountTransactorSession struct {
	Contract     *TradingAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TradingAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type TradingAccountRaw struct {
	Contract *TradingAccount // Generic contract binding to access the raw methods on
}

// TradingAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TradingAccountCallerRaw struct {
	Contract *TradingAccountCaller // Generic read-only contract binding to access the raw methods on
}

// TradingAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TradingAccountTransactorRaw struct {
	Contract *TradingAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTradingAccount creates a new instance of TradingAccount, bound to a specific deployed contract.
func NewTradingAccount(address common.Address, backend bind.ContractBackend) (*TradingAccount, error) {
	contract, err := bindTradingAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TradingAccount{TradingAccountCaller: TradingAccountCaller{contract: contract}, TradingAccountTransactor: TradingAccountTransactor{contract: contract}, TradingAccountFilterer: TradingAccountFilterer{contract: contract}}, nil
}

// NewTradingAccountCaller creates a new read-only instance of TradingAccount, bound to a specific deployed contract.
func NewTradingAccountCaller(address common.Address, caller bind.ContractCaller) (*TradingAccountCaller, error) {
	contract, err := bindTradingAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TradingAccountCaller{contract: contract}, nil
}

// NewTradingAccountTransactor creates a new write-only instance of TradingAccount, bound to a specific deployed contract.
func NewTradingAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*TradingAccountTransactor, error) {
	contract, err := bindTradingAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TradingAccountTransactor{contract: contract}, nil
}

// NewTradingAccountFilterer creates a new log filterer instance of TradingAccount, bound to a specific deployed contract.
func NewTradingAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*TradingAccountFilterer, error) {
	contract, err := bindTradingAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TradingAccountFilterer{contract: contract}, nil
}

// bindTradingAccount binds a generic wrapper to an already deployed contract.
func bindTradingAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TradingAccountABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TradingAccount *TradingAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TradingAccount.Contract.TradingAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TradingAccount *TradingAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradingAccount.Contract.TradingAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TradingAccount *TradingAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TradingAccount.Contract.TradingAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TradingAccount *TradingAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TradingAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TradingAccount *TradingAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradingAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TradingAccount *TradingAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TradingAccount.Contract.contract.Transact(opts, method, params...)
}

// CETH is a free data retrieval call binding the contract method 0xa1b4d011.
//
// Solidity: function cETH() view returns(address)
func (_TradingAccount *TradingAccountCaller) CETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TradingAccount.contract.Call(opts, &out, "cETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CETH is a free data retrieval call binding the contract method 0xa1b4d011.
//
// Solidity: function cETH() view returns(address)
func (_TradingAccount *TradingAccountSession) CETH() (common.Address, error) {
	return _TradingAccount.Contract.CETH(&_TradingAccount.CallOpts)
}

// CETH is a free data retrieval call binding the contract method 0xa1b4d011.
//
// Solidity: function cETH() view returns(address)
func (_TradingAccount *TradingAccountCallerSession) CETH() (common.Address, error) {
	return _TradingAccount.Contract.CETH(&_TradingAccount.CallOpts)
}

// CUSDC is a free data retrieval call binding the contract method 0xce347a3e.
//
// Solidity: function cUSDC() view returns(address)
func (_TradingAccount *TradingAccountCaller) CUSDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TradingAccount.contract.Call(opts, &out, "cUSDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CUSDC is a free data retrieval call binding the contract method 0xce347a3e.
//
// Solidity: function cUSDC() view returns(address)
func (_TradingAccount *TradingAccountSession) CUSDC() (common.Address, error) {
	return _TradingAccount.Contract.CUSDC(&_TradingAccount.CallOpts)
}

// CUSDC is a free data retrieval call binding the contract method 0xce347a3e.
//
// Solidity: function cUSDC() view returns(address)
func (_TradingAccount *TradingAccountCallerSession) CUSDC() (common.Address, error) {
	return _TradingAccount.Contract.CUSDC(&_TradingAccount.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_TradingAccount *TradingAccountCaller) Comptroller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TradingAccount.contract.Call(opts, &out, "comptroller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_TradingAccount *TradingAccountSession) Comptroller() (common.Address, error) {
	return _TradingAccount.Contract.Comptroller(&_TradingAccount.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_TradingAccount *TradingAccountCallerSession) Comptroller() (common.Address, error) {
	return _TradingAccount.Contract.Comptroller(&_TradingAccount.CallOpts)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0x8e15f473.
//
// Solidity: function getLatestPrice() view returns(uint256)
func (_TradingAccount *TradingAccountCaller) GetLatestPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TradingAccount.contract.Call(opts, &out, "getLatestPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPrice is a free data retrieval call binding the contract method 0x8e15f473.
//
// Solidity: function getLatestPrice() view returns(uint256)
func (_TradingAccount *TradingAccountSession) GetLatestPrice() (*big.Int, error) {
	return _TradingAccount.Contract.GetLatestPrice(&_TradingAccount.CallOpts)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0x8e15f473.
//
// Solidity: function getLatestPrice() view returns(uint256)
func (_TradingAccount *TradingAccountCallerSession) GetLatestPrice() (*big.Int, error) {
	return _TradingAccount.Contract.GetLatestPrice(&_TradingAccount.CallOpts)
}

// GetLimitOrder is a free data retrieval call binding the contract method 0x53061026.
//
// Solidity: function getLimitOrder(uint256 ) view returns(bool isLong, bool isCanceled, uint256 openSize, uint256 limitPrice, uint256 dealPrice, uint256 expireAt, address keeper)
func (_TradingAccount *TradingAccountCaller) GetLimitOrder(opts *bind.CallOpts, arg0 *big.Int) (struct {
	IsLong     bool
	IsCanceled bool
	OpenSize   *big.Int
	LimitPrice *big.Int
	DealPrice  *big.Int
	ExpireAt   *big.Int
	Keeper     common.Address
}, error) {
	var out []interface{}
	err := _TradingAccount.contract.Call(opts, &out, "getLimitOrder", arg0)

	outstruct := new(struct {
		IsLong     bool
		IsCanceled bool
		OpenSize   *big.Int
		LimitPrice *big.Int
		DealPrice  *big.Int
		ExpireAt   *big.Int
		Keeper     common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsLong = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.IsCanceled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.OpenSize = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LimitPrice = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.DealPrice = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.ExpireAt = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Keeper = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetLimitOrder is a free data retrieval call binding the contract method 0x53061026.
//
// Solidity: function getLimitOrder(uint256 ) view returns(bool isLong, bool isCanceled, uint256 openSize, uint256 limitPrice, uint256 dealPrice, uint256 expireAt, address keeper)
func (_TradingAccount *TradingAccountSession) GetLimitOrder(arg0 *big.Int) (struct {
	IsLong     bool
	IsCanceled bool
	OpenSize   *big.Int
	LimitPrice *big.Int
	DealPrice  *big.Int
	ExpireAt   *big.Int
	Keeper     common.Address
}, error) {
	return _TradingAccount.Contract.GetLimitOrder(&_TradingAccount.CallOpts, arg0)
}

// GetLimitOrder is a free data retrieval call binding the contract method 0x53061026.
//
// Solidity: function getLimitOrder(uint256 ) view returns(bool isLong, bool isCanceled, uint256 openSize, uint256 limitPrice, uint256 dealPrice, uint256 expireAt, address keeper)
func (_TradingAccount *TradingAccountCallerSession) GetLimitOrder(arg0 *big.Int) (struct {
	IsLong     bool
	IsCanceled bool
	OpenSize   *big.Int
	LimitPrice *big.Int
	DealPrice  *big.Int
	ExpireAt   *big.Int
	Keeper     common.Address
}, error) {
	return _TradingAccount.Contract.GetLimitOrder(&_TradingAccount.CallOpts, arg0)
}

// LastCEthBalance is a free data retrieval call binding the contract method 0x8fd1ef61.
//
// Solidity: function lastCEthBalance() view returns(uint256)
func (_TradingAccount *TradingAccountCaller) LastCEthBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TradingAccount.contract.Call(opts, &out, "lastCEthBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastCEthBalance is a free data retrieval call binding the contract method 0x8fd1ef61.
//
// Solidity: function lastCEthBalance() view returns(uint256)
func (_TradingAccount *TradingAccountSession) LastCEthBalance() (*big.Int, error) {
	return _TradingAccount.Contract.LastCEthBalance(&_TradingAccount.CallOpts)
}

// LastCEthBalance is a free data retrieval call binding the contract method 0x8fd1ef61.
//
// Solidity: function lastCEthBalance() view returns(uint256)
func (_TradingAccount *TradingAccountCallerSession) LastCEthBalance() (*big.Int, error) {
	return _TradingAccount.Contract.LastCEthBalance(&_TradingAccount.CallOpts)
}

// LastCUsdcBalance is a free data retrieval call binding the contract method 0x883f9741.
//
// Solidity: function lastCUsdcBalance() view returns(uint256)
func (_TradingAccount *TradingAccountCaller) LastCUsdcBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TradingAccount.contract.Call(opts, &out, "lastCUsdcBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastCUsdcBalance is a free data retrieval call binding the contract method 0x883f9741.
//
// Solidity: function lastCUsdcBalance() view returns(uint256)
func (_TradingAccount *TradingAccountSession) LastCUsdcBalance() (*big.Int, error) {
	return _TradingAccount.Contract.LastCUsdcBalance(&_TradingAccount.CallOpts)
}

// LastCUsdcBalance is a free data retrieval call binding the contract method 0x883f9741.
//
// Solidity: function lastCUsdcBalance() view returns(uint256)
func (_TradingAccount *TradingAccountCallerSession) LastCUsdcBalance() (*big.Int, error) {
	return _TradingAccount.Contract.LastCUsdcBalance(&_TradingAccount.CallOpts)
}

// LastOrderId is a free data retrieval call binding the contract method 0x5662ecc7.
//
// Solidity: function lastOrderId() view returns(uint256)
func (_TradingAccount *TradingAccountCaller) LastOrderId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TradingAccount.contract.Call(opts, &out, "lastOrderId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastOrderId is a free data retrieval call binding the contract method 0x5662ecc7.
//
// Solidity: function lastOrderId() view returns(uint256)
func (_TradingAccount *TradingAccountSession) LastOrderId() (*big.Int, error) {
	return _TradingAccount.Contract.LastOrderId(&_TradingAccount.CallOpts)
}

// LastOrderId is a free data retrieval call binding the contract method 0x5662ecc7.
//
// Solidity: function lastOrderId() view returns(uint256)
func (_TradingAccount *TradingAccountCallerSession) LastOrderId() (*big.Int, error) {
	return _TradingAccount.Contract.LastOrderId(&_TradingAccount.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TradingAccount *TradingAccountCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TradingAccount.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TradingAccount *TradingAccountSession) Owner() (common.Address, error) {
	return _TradingAccount.Contract.Owner(&_TradingAccount.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TradingAccount *TradingAccountCallerSession) Owner() (common.Address, error) {
	return _TradingAccount.Contract.Owner(&_TradingAccount.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_TradingAccount *TradingAccountCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TradingAccount.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_TradingAccount *TradingAccountSession) PendingOwner() (common.Address, error) {
	return _TradingAccount.Contract.PendingOwner(&_TradingAccount.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_TradingAccount *TradingAccountCallerSession) PendingOwner() (common.Address, error) {
	return _TradingAccount.Contract.PendingOwner(&_TradingAccount.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_TradingAccount *TradingAccountCaller) PriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TradingAccount.contract.Call(opts, &out, "priceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_TradingAccount *TradingAccountSession) PriceFeed() (common.Address, error) {
	return _TradingAccount.Contract.PriceFeed(&_TradingAccount.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_TradingAccount *TradingAccountCallerSession) PriceFeed() (common.Address, error) {
	return _TradingAccount.Contract.PriceFeed(&_TradingAccount.CallOpts)
}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_TradingAccount *TradingAccountCaller) UniswapV2Pair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TradingAccount.contract.Call(opts, &out, "uniswapV2Pair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_TradingAccount *TradingAccountSession) UniswapV2Pair() (common.Address, error) {
	return _TradingAccount.Contract.UniswapV2Pair(&_TradingAccount.CallOpts)
}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_TradingAccount *TradingAccountCallerSession) UniswapV2Pair() (common.Address, error) {
	return _TradingAccount.Contract.UniswapV2Pair(&_TradingAccount.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_TradingAccount *TradingAccountCaller) Usdc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TradingAccount.contract.Call(opts, &out, "usdc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_TradingAccount *TradingAccountSession) Usdc() (common.Address, error) {
	return _TradingAccount.Contract.Usdc(&_TradingAccount.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_TradingAccount *TradingAccountCallerSession) Usdc() (common.Address, error) {
	return _TradingAccount.Contract.Usdc(&_TradingAccount.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_TradingAccount *TradingAccountCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TradingAccount.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_TradingAccount *TradingAccountSession) Weth() (common.Address, error) {
	return _TradingAccount.Contract.Weth(&_TradingAccount.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_TradingAccount *TradingAccountCallerSession) Weth() (common.Address, error) {
	return _TradingAccount.Contract.Weth(&_TradingAccount.CallOpts)
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xebbc4965.
//
// Solidity: function acceptOwner() returns()
func (_TradingAccount *TradingAccountTransactor) AcceptOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradingAccount.contract.Transact(opts, "acceptOwner")
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xebbc4965.
//
// Solidity: function acceptOwner() returns()
func (_TradingAccount *TradingAccountSession) AcceptOwner() (*types.Transaction, error) {
	return _TradingAccount.Contract.AcceptOwner(&_TradingAccount.TransactOpts)
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xebbc4965.
//
// Solidity: function acceptOwner() returns()
func (_TradingAccount *TradingAccountTransactorSession) AcceptOwner() (*types.Transaction, error) {
	return _TradingAccount.Contract.AcceptOwner(&_TradingAccount.TransactOpts)
}

// CancelLimitOrder is a paid mutator transaction binding the contract method 0xa5cdc8fc.
//
// Solidity: function cancelLimitOrder(uint256 orderId) returns()
func (_TradingAccount *TradingAccountTransactor) CancelLimitOrder(opts *bind.TransactOpts, orderId *big.Int) (*types.Transaction, error) {
	return _TradingAccount.contract.Transact(opts, "cancelLimitOrder", orderId)
}

// CancelLimitOrder is a paid mutator transaction binding the contract method 0xa5cdc8fc.
//
// Solidity: function cancelLimitOrder(uint256 orderId) returns()
func (_TradingAccount *TradingAccountSession) CancelLimitOrder(orderId *big.Int) (*types.Transaction, error) {
	return _TradingAccount.Contract.CancelLimitOrder(&_TradingAccount.TransactOpts, orderId)
}

// CancelLimitOrder is a paid mutator transaction binding the contract method 0xa5cdc8fc.
//
// Solidity: function cancelLimitOrder(uint256 orderId) returns()
func (_TradingAccount *TradingAccountTransactorSession) CancelLimitOrder(orderId *big.Int) (*types.Transaction, error) {
	return _TradingAccount.Contract.CancelLimitOrder(&_TradingAccount.TransactOpts, orderId)
}

// CloseLong is a paid mutator transaction binding the contract method 0x6d89ba2a.
//
// Solidity: function closeLong(uint256 usdcAmount, bool closeAll) returns(uint256 repayAmount)
func (_TradingAccount *TradingAccountTransactor) CloseLong(opts *bind.TransactOpts, usdcAmount *big.Int, closeAll bool) (*types.Transaction, error) {
	return _TradingAccount.contract.Transact(opts, "closeLong", usdcAmount, closeAll)
}

// CloseLong is a paid mutator transaction binding the contract method 0x6d89ba2a.
//
// Solidity: function closeLong(uint256 usdcAmount, bool closeAll) returns(uint256 repayAmount)
func (_TradingAccount *TradingAccountSession) CloseLong(usdcAmount *big.Int, closeAll bool) (*types.Transaction, error) {
	return _TradingAccount.Contract.CloseLong(&_TradingAccount.TransactOpts, usdcAmount, closeAll)
}

// CloseLong is a paid mutator transaction binding the contract method 0x6d89ba2a.
//
// Solidity: function closeLong(uint256 usdcAmount, bool closeAll) returns(uint256 repayAmount)
func (_TradingAccount *TradingAccountTransactorSession) CloseLong(usdcAmount *big.Int, closeAll bool) (*types.Transaction, error) {
	return _TradingAccount.Contract.CloseLong(&_TradingAccount.TransactOpts, usdcAmount, closeAll)
}

// CloseShort is a paid mutator transaction binding the contract method 0xf083ff9a.
//
// Solidity: function closeShort(uint256 ethAmount, bool closeAll) returns(uint256 repayAmount)
func (_TradingAccount *TradingAccountTransactor) CloseShort(opts *bind.TransactOpts, ethAmount *big.Int, closeAll bool) (*types.Transaction, error) {
	return _TradingAccount.contract.Transact(opts, "closeShort", ethAmount, closeAll)
}

// CloseShort is a paid mutator transaction binding the contract method 0xf083ff9a.
//
// Solidity: function closeShort(uint256 ethAmount, bool closeAll) returns(uint256 repayAmount)
func (_TradingAccount *TradingAccountSession) CloseShort(ethAmount *big.Int, closeAll bool) (*types.Transaction, error) {
	return _TradingAccount.Contract.CloseShort(&_TradingAccount.TransactOpts, ethAmount, closeAll)
}

// CloseShort is a paid mutator transaction binding the contract method 0xf083ff9a.
//
// Solidity: function closeShort(uint256 ethAmount, bool closeAll) returns(uint256 repayAmount)
func (_TradingAccount *TradingAccountTransactorSession) CloseShort(ethAmount *big.Int, closeAll bool) (*types.Transaction, error) {
	return _TradingAccount.Contract.CloseShort(&_TradingAccount.TransactOpts, ethAmount, closeAll)
}

// DepositETH is a paid mutator transaction binding the contract method 0xf6326fb3.
//
// Solidity: function depositETH() payable returns()
func (_TradingAccount *TradingAccountTransactor) DepositETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradingAccount.contract.Transact(opts, "depositETH")
}

// DepositETH is a paid mutator transaction binding the contract method 0xf6326fb3.
//
// Solidity: function depositETH() payable returns()
func (_TradingAccount *TradingAccountSession) DepositETH() (*types.Transaction, error) {
	return _TradingAccount.Contract.DepositETH(&_TradingAccount.TransactOpts)
}

// DepositETH is a paid mutator transaction binding the contract method 0xf6326fb3.
//
// Solidity: function depositETH() payable returns()
func (_TradingAccount *TradingAccountTransactorSession) DepositETH() (*types.Transaction, error) {
	return _TradingAccount.Contract.DepositETH(&_TradingAccount.TransactOpts)
}

// DepositUSDC is a paid mutator transaction binding the contract method 0xf688bcfb.
//
// Solidity: function depositUSDC(uint256 amount) returns()
func (_TradingAccount *TradingAccountTransactor) DepositUSDC(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TradingAccount.contract.Transact(opts, "depositUSDC", amount)
}

// DepositUSDC is a paid mutator transaction binding the contract method 0xf688bcfb.
//
// Solidity: function depositUSDC(uint256 amount) returns()
func (_TradingAccount *TradingAccountSession) DepositUSDC(amount *big.Int) (*types.Transaction, error) {
	return _TradingAccount.Contract.DepositUSDC(&_TradingAccount.TransactOpts, amount)
}

// DepositUSDC is a paid mutator transaction binding the contract method 0xf688bcfb.
//
// Solidity: function depositUSDC(uint256 amount) returns()
func (_TradingAccount *TradingAccountTransactorSession) DepositUSDC(amount *big.Int) (*types.Transaction, error) {
	return _TradingAccount.Contract.DepositUSDC(&_TradingAccount.TransactOpts, amount)
}

// ExecuteLimitOrder is a paid mutator transaction binding the contract method 0xc7219460.
//
// Solidity: function executeLimitOrder(uint256 orderId) returns()
func (_TradingAccount *TradingAccountTransactor) ExecuteLimitOrder(opts *bind.TransactOpts, orderId *big.Int) (*types.Transaction, error) {
	return _TradingAccount.contract.Transact(opts, "executeLimitOrder", orderId)
}

// ExecuteLimitOrder is a paid mutator transaction binding the contract method 0xc7219460.
//
// Solidity: function executeLimitOrder(uint256 orderId) returns()
func (_TradingAccount *TradingAccountSession) ExecuteLimitOrder(orderId *big.Int) (*types.Transaction, error) {
	return _TradingAccount.Contract.ExecuteLimitOrder(&_TradingAccount.TransactOpts, orderId)
}

// ExecuteLimitOrder is a paid mutator transaction binding the contract method 0xc7219460.
//
// Solidity: function executeLimitOrder(uint256 orderId) returns()
func (_TradingAccount *TradingAccountTransactorSession) ExecuteLimitOrder(orderId *big.Int) (*types.Transaction, error) {
	return _TradingAccount.Contract.ExecuteLimitOrder(&_TradingAccount.TransactOpts, orderId)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a29e2de.
//
// Solidity: function initialize(address owner_, address weth_, address usdc_, address uniswapV2Pair_, address cETH_, address cUSDC_, address comptroller_, address priceFeed_) returns()
func (_TradingAccount *TradingAccountTransactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, weth_ common.Address, usdc_ common.Address, uniswapV2Pair_ common.Address, cETH_ common.Address, cUSDC_ common.Address, comptroller_ common.Address, priceFeed_ common.Address) (*types.Transaction, error) {
	return _TradingAccount.contract.Transact(opts, "initialize", owner_, weth_, usdc_, uniswapV2Pair_, cETH_, cUSDC_, comptroller_, priceFeed_)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a29e2de.
//
// Solidity: function initialize(address owner_, address weth_, address usdc_, address uniswapV2Pair_, address cETH_, address cUSDC_, address comptroller_, address priceFeed_) returns()
func (_TradingAccount *TradingAccountSession) Initialize(owner_ common.Address, weth_ common.Address, usdc_ common.Address, uniswapV2Pair_ common.Address, cETH_ common.Address, cUSDC_ common.Address, comptroller_ common.Address, priceFeed_ common.Address) (*types.Transaction, error) {
	return _TradingAccount.Contract.Initialize(&_TradingAccount.TransactOpts, owner_, weth_, usdc_, uniswapV2Pair_, cETH_, cUSDC_, comptroller_, priceFeed_)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a29e2de.
//
// Solidity: function initialize(address owner_, address weth_, address usdc_, address uniswapV2Pair_, address cETH_, address cUSDC_, address comptroller_, address priceFeed_) returns()
func (_TradingAccount *TradingAccountTransactorSession) Initialize(owner_ common.Address, weth_ common.Address, usdc_ common.Address, uniswapV2Pair_ common.Address, cETH_ common.Address, cUSDC_ common.Address, comptroller_ common.Address, priceFeed_ common.Address) (*types.Transaction, error) {
	return _TradingAccount.Contract.Initialize(&_TradingAccount.TransactOpts, owner_, weth_, usdc_, uniswapV2Pair_, cETH_, cUSDC_, comptroller_, priceFeed_)
}

// LimitOpenLong is a paid mutator transaction binding the contract method 0x09fffa84.
//
// Solidity: function limitOpenLong(uint256 ethSize, uint256 limitPrice, uint256 expireAt, address keeper) returns(uint256 orderId)
func (_TradingAccount *TradingAccountTransactor) LimitOpenLong(opts *bind.TransactOpts, ethSize *big.Int, limitPrice *big.Int, expireAt *big.Int, keeper common.Address) (*types.Transaction, error) {
	return _TradingAccount.contract.Transact(opts, "limitOpenLong", ethSize, limitPrice, expireAt, keeper)
}

// LimitOpenLong is a paid mutator transaction binding the contract method 0x09fffa84.
//
// Solidity: function limitOpenLong(uint256 ethSize, uint256 limitPrice, uint256 expireAt, address keeper) returns(uint256 orderId)
func (_TradingAccount *TradingAccountSession) LimitOpenLong(ethSize *big.Int, limitPrice *big.Int, expireAt *big.Int, keeper common.Address) (*types.Transaction, error) {
	return _TradingAccount.Contract.LimitOpenLong(&_TradingAccount.TransactOpts, ethSize, limitPrice, expireAt, keeper)
}

// LimitOpenLong is a paid mutator transaction binding the contract method 0x09fffa84.
//
// Solidity: function limitOpenLong(uint256 ethSize, uint256 limitPrice, uint256 expireAt, address keeper) returns(uint256 orderId)
func (_TradingAccount *TradingAccountTransactorSession) LimitOpenLong(ethSize *big.Int, limitPrice *big.Int, expireAt *big.Int, keeper common.Address) (*types.Transaction, error) {
	return _TradingAccount.Contract.LimitOpenLong(&_TradingAccount.TransactOpts, ethSize, limitPrice, expireAt, keeper)
}

// LimitOpenShort is a paid mutator transaction binding the contract method 0x3b7ba770.
//
// Solidity: function limitOpenShort(uint256 usdcSize, uint256 limitPrice, uint256 expireAt, address keeper) returns(uint256 orderId)
func (_TradingAccount *TradingAccountTransactor) LimitOpenShort(opts *bind.TransactOpts, usdcSize *big.Int, limitPrice *big.Int, expireAt *big.Int, keeper common.Address) (*types.Transaction, error) {
	return _TradingAccount.contract.Transact(opts, "limitOpenShort", usdcSize, limitPrice, expireAt, keeper)
}

// LimitOpenShort is a paid mutator transaction binding the contract method 0x3b7ba770.
//
// Solidity: function limitOpenShort(uint256 usdcSize, uint256 limitPrice, uint256 expireAt, address keeper) returns(uint256 orderId)
func (_TradingAccount *TradingAccountSession) LimitOpenShort(usdcSize *big.Int, limitPrice *big.Int, expireAt *big.Int, keeper common.Address) (*types.Transaction, error) {
	return _TradingAccount.Contract.LimitOpenShort(&_TradingAccount.TransactOpts, usdcSize, limitPrice, expireAt, keeper)
}

// LimitOpenShort is a paid mutator transaction binding the contract method 0x3b7ba770.
//
// Solidity: function limitOpenShort(uint256 usdcSize, uint256 limitPrice, uint256 expireAt, address keeper) returns(uint256 orderId)
func (_TradingAccount *TradingAccountTransactorSession) LimitOpenShort(usdcSize *big.Int, limitPrice *big.Int, expireAt *big.Int, keeper common.Address) (*types.Transaction, error) {
	return _TradingAccount.Contract.LimitOpenShort(&_TradingAccount.TransactOpts, usdcSize, limitPrice, expireAt, keeper)
}

// OpenLong is a paid mutator transaction binding the contract method 0x9f039090.
//
// Solidity: function openLong(uint256 ethSize) returns()
func (_TradingAccount *TradingAccountTransactor) OpenLong(opts *bind.TransactOpts, ethSize *big.Int) (*types.Transaction, error) {
	return _TradingAccount.contract.Transact(opts, "openLong", ethSize)
}

// OpenLong is a paid mutator transaction binding the contract method 0x9f039090.
//
// Solidity: function openLong(uint256 ethSize) returns()
func (_TradingAccount *TradingAccountSession) OpenLong(ethSize *big.Int) (*types.Transaction, error) {
	return _TradingAccount.Contract.OpenLong(&_TradingAccount.TransactOpts, ethSize)
}

// OpenLong is a paid mutator transaction binding the contract method 0x9f039090.
//
// Solidity: function openLong(uint256 ethSize) returns()
func (_TradingAccount *TradingAccountTransactorSession) OpenLong(ethSize *big.Int) (*types.Transaction, error) {
	return _TradingAccount.Contract.OpenLong(&_TradingAccount.TransactOpts, ethSize)
}

// OpenShort is a paid mutator transaction binding the contract method 0xed072e22.
//
// Solidity: function openShort(uint256 usdcSize) returns()
func (_TradingAccount *TradingAccountTransactor) OpenShort(opts *bind.TransactOpts, usdcSize *big.Int) (*types.Transaction, error) {
	return _TradingAccount.contract.Transact(opts, "openShort", usdcSize)
}

// OpenShort is a paid mutator transaction binding the contract method 0xed072e22.
//
// Solidity: function openShort(uint256 usdcSize) returns()
func (_TradingAccount *TradingAccountSession) OpenShort(usdcSize *big.Int) (*types.Transaction, error) {
	return _TradingAccount.Contract.OpenShort(&_TradingAccount.TransactOpts, usdcSize)
}

// OpenShort is a paid mutator transaction binding the contract method 0xed072e22.
//
// Solidity: function openShort(uint256 usdcSize) returns()
func (_TradingAccount *TradingAccountTransactorSession) OpenShort(usdcSize *big.Int) (*types.Transaction, error) {
	return _TradingAccount.Contract.OpenShort(&_TradingAccount.TransactOpts, usdcSize)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address newPendingOwner) returns()
func (_TradingAccount *TradingAccountTransactor) SetPendingOwner(opts *bind.TransactOpts, newPendingOwner common.Address) (*types.Transaction, error) {
	return _TradingAccount.contract.Transact(opts, "setPendingOwner", newPendingOwner)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address newPendingOwner) returns()
func (_TradingAccount *TradingAccountSession) SetPendingOwner(newPendingOwner common.Address) (*types.Transaction, error) {
	return _TradingAccount.Contract.SetPendingOwner(&_TradingAccount.TransactOpts, newPendingOwner)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address newPendingOwner) returns()
func (_TradingAccount *TradingAccountTransactorSession) SetPendingOwner(newPendingOwner common.Address) (*types.Transaction, error) {
	return _TradingAccount.Contract.SetPendingOwner(&_TradingAccount.TransactOpts, newPendingOwner)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_TradingAccount *TradingAccountTransactor) UniswapV2Call(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _TradingAccount.contract.Transact(opts, "uniswapV2Call", sender, amount0, amount1, data)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_TradingAccount *TradingAccountSession) UniswapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _TradingAccount.Contract.UniswapV2Call(&_TradingAccount.TransactOpts, sender, amount0, amount1, data)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_TradingAccount *TradingAccountTransactorSession) UniswapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _TradingAccount.Contract.UniswapV2Call(&_TradingAccount.TransactOpts, sender, amount0, amount1, data)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xc7cdea37.
//
// Solidity: function withdrawETH(uint256 cEthAmount, uint256 ethAmount) returns()
func (_TradingAccount *TradingAccountTransactor) WithdrawETH(opts *bind.TransactOpts, cEthAmount *big.Int, ethAmount *big.Int) (*types.Transaction, error) {
	return _TradingAccount.contract.Transact(opts, "withdrawETH", cEthAmount, ethAmount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xc7cdea37.
//
// Solidity: function withdrawETH(uint256 cEthAmount, uint256 ethAmount) returns()
func (_TradingAccount *TradingAccountSession) WithdrawETH(cEthAmount *big.Int, ethAmount *big.Int) (*types.Transaction, error) {
	return _TradingAccount.Contract.WithdrawETH(&_TradingAccount.TransactOpts, cEthAmount, ethAmount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xc7cdea37.
//
// Solidity: function withdrawETH(uint256 cEthAmount, uint256 ethAmount) returns()
func (_TradingAccount *TradingAccountTransactorSession) WithdrawETH(cEthAmount *big.Int, ethAmount *big.Int) (*types.Transaction, error) {
	return _TradingAccount.Contract.WithdrawETH(&_TradingAccount.TransactOpts, cEthAmount, ethAmount)
}

// WithdrawUSDC is a paid mutator transaction binding the contract method 0x88e08665.
//
// Solidity: function withdrawUSDC(uint256 cUsdcAmount, uint256 usdcAmount) returns()
func (_TradingAccount *TradingAccountTransactor) WithdrawUSDC(opts *bind.TransactOpts, cUsdcAmount *big.Int, usdcAmount *big.Int) (*types.Transaction, error) {
	return _TradingAccount.contract.Transact(opts, "withdrawUSDC", cUsdcAmount, usdcAmount)
}

// WithdrawUSDC is a paid mutator transaction binding the contract method 0x88e08665.
//
// Solidity: function withdrawUSDC(uint256 cUsdcAmount, uint256 usdcAmount) returns()
func (_TradingAccount *TradingAccountSession) WithdrawUSDC(cUsdcAmount *big.Int, usdcAmount *big.Int) (*types.Transaction, error) {
	return _TradingAccount.Contract.WithdrawUSDC(&_TradingAccount.TransactOpts, cUsdcAmount, usdcAmount)
}

// WithdrawUSDC is a paid mutator transaction binding the contract method 0x88e08665.
//
// Solidity: function withdrawUSDC(uint256 cUsdcAmount, uint256 usdcAmount) returns()
func (_TradingAccount *TradingAccountTransactorSession) WithdrawUSDC(cUsdcAmount *big.Int, usdcAmount *big.Int) (*types.Transaction, error) {
	return _TradingAccount.Contract.WithdrawUSDC(&_TradingAccount.TransactOpts, cUsdcAmount, usdcAmount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TradingAccount *TradingAccountTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradingAccount.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TradingAccount *TradingAccountSession) Receive() (*types.Transaction, error) {
	return _TradingAccount.Contract.Receive(&_TradingAccount.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TradingAccount *TradingAccountTransactorSession) Receive() (*types.Transaction, error) {
	return _TradingAccount.Contract.Receive(&_TradingAccount.TransactOpts)
}

// TradingAccountCancelLimitOrderIterator is returned from FilterCancelLimitOrder and is used to iterate over the raw logs and unpacked data for CancelLimitOrder events raised by the TradingAccount contract.
type TradingAccountCancelLimitOrderIterator struct {
	Event *TradingAccountCancelLimitOrder // Event containing the contract specifics and raw log

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
func (it *TradingAccountCancelLimitOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingAccountCancelLimitOrder)
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
		it.Event = new(TradingAccountCancelLimitOrder)
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
func (it *TradingAccountCancelLimitOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingAccountCancelLimitOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingAccountCancelLimitOrder represents a CancelLimitOrder event raised by the TradingAccount contract.
type TradingAccountCancelLimitOrder struct {
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCancelLimitOrder is a free log retrieval operation binding the contract event 0xacb097d265630549c66967ff31a6d9a30a35656b951b1fe2d90a3499ba29ee26.
//
// Solidity: event CancelLimitOrder(uint256 orderId)
func (_TradingAccount *TradingAccountFilterer) FilterCancelLimitOrder(opts *bind.FilterOpts) (*TradingAccountCancelLimitOrderIterator, error) {

	logs, sub, err := _TradingAccount.contract.FilterLogs(opts, "CancelLimitOrder")
	if err != nil {
		return nil, err
	}
	return &TradingAccountCancelLimitOrderIterator{contract: _TradingAccount.contract, event: "CancelLimitOrder", logs: logs, sub: sub}, nil
}

// WatchCancelLimitOrder is a free log subscription operation binding the contract event 0xacb097d265630549c66967ff31a6d9a30a35656b951b1fe2d90a3499ba29ee26.
//
// Solidity: event CancelLimitOrder(uint256 orderId)
func (_TradingAccount *TradingAccountFilterer) WatchCancelLimitOrder(opts *bind.WatchOpts, sink chan<- *TradingAccountCancelLimitOrder) (event.Subscription, error) {

	logs, sub, err := _TradingAccount.contract.WatchLogs(opts, "CancelLimitOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingAccountCancelLimitOrder)
				if err := _TradingAccount.contract.UnpackLog(event, "CancelLimitOrder", log); err != nil {
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

// ParseCancelLimitOrder is a log parse operation binding the contract event 0xacb097d265630549c66967ff31a6d9a30a35656b951b1fe2d90a3499ba29ee26.
//
// Solidity: event CancelLimitOrder(uint256 orderId)
func (_TradingAccount *TradingAccountFilterer) ParseCancelLimitOrder(log types.Log) (*TradingAccountCancelLimitOrder, error) {
	event := new(TradingAccountCancelLimitOrder)
	if err := _TradingAccount.contract.UnpackLog(event, "CancelLimitOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingAccountCloseLongIterator is returned from FilterCloseLong and is used to iterate over the raw logs and unpacked data for CloseLong events raised by the TradingAccount contract.
type TradingAccountCloseLongIterator struct {
	Event *TradingAccountCloseLong // Event containing the contract specifics and raw log

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
func (it *TradingAccountCloseLongIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingAccountCloseLong)
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
		it.Event = new(TradingAccountCloseLong)
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
func (it *TradingAccountCloseLongIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingAccountCloseLongIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingAccountCloseLong represents a CloseLong event raised by the TradingAccount contract.
type TradingAccountCloseLong struct {
	RedeemCTokenAmount     *big.Int
	RedeemUnderlyingAmount *big.Int
	RepayAmount            *big.Int
	CloseAll               bool
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterCloseLong is a free log retrieval operation binding the contract event 0x37b43c784f8d7bce65eda84311e0e462cf6fd3db2f40193ef86e97b610f0c06f.
//
// Solidity: event CloseLong(uint256 redeemCTokenAmount, uint256 redeemUnderlyingAmount, uint256 repayAmount, bool closeAll)
func (_TradingAccount *TradingAccountFilterer) FilterCloseLong(opts *bind.FilterOpts) (*TradingAccountCloseLongIterator, error) {

	logs, sub, err := _TradingAccount.contract.FilterLogs(opts, "CloseLong")
	if err != nil {
		return nil, err
	}
	return &TradingAccountCloseLongIterator{contract: _TradingAccount.contract, event: "CloseLong", logs: logs, sub: sub}, nil
}

// WatchCloseLong is a free log subscription operation binding the contract event 0x37b43c784f8d7bce65eda84311e0e462cf6fd3db2f40193ef86e97b610f0c06f.
//
// Solidity: event CloseLong(uint256 redeemCTokenAmount, uint256 redeemUnderlyingAmount, uint256 repayAmount, bool closeAll)
func (_TradingAccount *TradingAccountFilterer) WatchCloseLong(opts *bind.WatchOpts, sink chan<- *TradingAccountCloseLong) (event.Subscription, error) {

	logs, sub, err := _TradingAccount.contract.WatchLogs(opts, "CloseLong")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingAccountCloseLong)
				if err := _TradingAccount.contract.UnpackLog(event, "CloseLong", log); err != nil {
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

// ParseCloseLong is a log parse operation binding the contract event 0x37b43c784f8d7bce65eda84311e0e462cf6fd3db2f40193ef86e97b610f0c06f.
//
// Solidity: event CloseLong(uint256 redeemCTokenAmount, uint256 redeemUnderlyingAmount, uint256 repayAmount, bool closeAll)
func (_TradingAccount *TradingAccountFilterer) ParseCloseLong(log types.Log) (*TradingAccountCloseLong, error) {
	event := new(TradingAccountCloseLong)
	if err := _TradingAccount.contract.UnpackLog(event, "CloseLong", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingAccountCloseShortIterator is returned from FilterCloseShort and is used to iterate over the raw logs and unpacked data for CloseShort events raised by the TradingAccount contract.
type TradingAccountCloseShortIterator struct {
	Event *TradingAccountCloseShort // Event containing the contract specifics and raw log

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
func (it *TradingAccountCloseShortIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingAccountCloseShort)
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
		it.Event = new(TradingAccountCloseShort)
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
func (it *TradingAccountCloseShortIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingAccountCloseShortIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingAccountCloseShort represents a CloseShort event raised by the TradingAccount contract.
type TradingAccountCloseShort struct {
	RedeemCTokenAmount     *big.Int
	RedeemUnderlyingAmount *big.Int
	RepayAmount            *big.Int
	CloseAll               bool
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterCloseShort is a free log retrieval operation binding the contract event 0x78ce03753721a9db485dbe4101e71eb7c530f35bf14365d3b53e0465bc54c0c5.
//
// Solidity: event CloseShort(uint256 redeemCTokenAmount, uint256 redeemUnderlyingAmount, uint256 repayAmount, bool closeAll)
func (_TradingAccount *TradingAccountFilterer) FilterCloseShort(opts *bind.FilterOpts) (*TradingAccountCloseShortIterator, error) {

	logs, sub, err := _TradingAccount.contract.FilterLogs(opts, "CloseShort")
	if err != nil {
		return nil, err
	}
	return &TradingAccountCloseShortIterator{contract: _TradingAccount.contract, event: "CloseShort", logs: logs, sub: sub}, nil
}

// WatchCloseShort is a free log subscription operation binding the contract event 0x78ce03753721a9db485dbe4101e71eb7c530f35bf14365d3b53e0465bc54c0c5.
//
// Solidity: event CloseShort(uint256 redeemCTokenAmount, uint256 redeemUnderlyingAmount, uint256 repayAmount, bool closeAll)
func (_TradingAccount *TradingAccountFilterer) WatchCloseShort(opts *bind.WatchOpts, sink chan<- *TradingAccountCloseShort) (event.Subscription, error) {

	logs, sub, err := _TradingAccount.contract.WatchLogs(opts, "CloseShort")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingAccountCloseShort)
				if err := _TradingAccount.contract.UnpackLog(event, "CloseShort", log); err != nil {
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

// ParseCloseShort is a log parse operation binding the contract event 0x78ce03753721a9db485dbe4101e71eb7c530f35bf14365d3b53e0465bc54c0c5.
//
// Solidity: event CloseShort(uint256 redeemCTokenAmount, uint256 redeemUnderlyingAmount, uint256 repayAmount, bool closeAll)
func (_TradingAccount *TradingAccountFilterer) ParseCloseShort(log types.Log) (*TradingAccountCloseShort, error) {
	event := new(TradingAccountCloseShort)
	if err := _TradingAccount.contract.UnpackLog(event, "CloseShort", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingAccountDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the TradingAccount contract.
type TradingAccountDepositIterator struct {
	Event *TradingAccountDeposit // Event containing the contract specifics and raw log

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
func (it *TradingAccountDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingAccountDeposit)
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
		it.Event = new(TradingAccountDeposit)
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
func (it *TradingAccountDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingAccountDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingAccountDeposit represents a Deposit event raised by the TradingAccount contract.
type TradingAccountDeposit struct {
	IsETH        bool
	Amount       *big.Int
	CTokenAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x1c49111414d8cd618ba259b16508c34541a4f27a1e3c4372c0f358d5086652bd.
//
// Solidity: event Deposit(bool isETH, uint256 amount, uint256 cTokenAmount)
func (_TradingAccount *TradingAccountFilterer) FilterDeposit(opts *bind.FilterOpts) (*TradingAccountDepositIterator, error) {

	logs, sub, err := _TradingAccount.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &TradingAccountDepositIterator{contract: _TradingAccount.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x1c49111414d8cd618ba259b16508c34541a4f27a1e3c4372c0f358d5086652bd.
//
// Solidity: event Deposit(bool isETH, uint256 amount, uint256 cTokenAmount)
func (_TradingAccount *TradingAccountFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *TradingAccountDeposit) (event.Subscription, error) {

	logs, sub, err := _TradingAccount.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingAccountDeposit)
				if err := _TradingAccount.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x1c49111414d8cd618ba259b16508c34541a4f27a1e3c4372c0f358d5086652bd.
//
// Solidity: event Deposit(bool isETH, uint256 amount, uint256 cTokenAmount)
func (_TradingAccount *TradingAccountFilterer) ParseDeposit(log types.Log) (*TradingAccountDeposit, error) {
	event := new(TradingAccountDeposit)
	if err := _TradingAccount.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingAccountExecuteLimitOrderIterator is returned from FilterExecuteLimitOrder and is used to iterate over the raw logs and unpacked data for ExecuteLimitOrder events raised by the TradingAccount contract.
type TradingAccountExecuteLimitOrderIterator struct {
	Event *TradingAccountExecuteLimitOrder // Event containing the contract specifics and raw log

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
func (it *TradingAccountExecuteLimitOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingAccountExecuteLimitOrder)
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
		it.Event = new(TradingAccountExecuteLimitOrder)
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
func (it *TradingAccountExecuteLimitOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingAccountExecuteLimitOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingAccountExecuteLimitOrder represents a ExecuteLimitOrder event raised by the TradingAccount contract.
type TradingAccountExecuteLimitOrder struct {
	OrderId   *big.Int
	DealPrice *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExecuteLimitOrder is a free log retrieval operation binding the contract event 0xba5c68549d7afe892d9d1e6d4acdb167b18828851f715d44db58adb8b0f91aca.
//
// Solidity: event ExecuteLimitOrder(uint256 orderId, uint256 dealPrice)
func (_TradingAccount *TradingAccountFilterer) FilterExecuteLimitOrder(opts *bind.FilterOpts) (*TradingAccountExecuteLimitOrderIterator, error) {

	logs, sub, err := _TradingAccount.contract.FilterLogs(opts, "ExecuteLimitOrder")
	if err != nil {
		return nil, err
	}
	return &TradingAccountExecuteLimitOrderIterator{contract: _TradingAccount.contract, event: "ExecuteLimitOrder", logs: logs, sub: sub}, nil
}

// WatchExecuteLimitOrder is a free log subscription operation binding the contract event 0xba5c68549d7afe892d9d1e6d4acdb167b18828851f715d44db58adb8b0f91aca.
//
// Solidity: event ExecuteLimitOrder(uint256 orderId, uint256 dealPrice)
func (_TradingAccount *TradingAccountFilterer) WatchExecuteLimitOrder(opts *bind.WatchOpts, sink chan<- *TradingAccountExecuteLimitOrder) (event.Subscription, error) {

	logs, sub, err := _TradingAccount.contract.WatchLogs(opts, "ExecuteLimitOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingAccountExecuteLimitOrder)
				if err := _TradingAccount.contract.UnpackLog(event, "ExecuteLimitOrder", log); err != nil {
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

// ParseExecuteLimitOrder is a log parse operation binding the contract event 0xba5c68549d7afe892d9d1e6d4acdb167b18828851f715d44db58adb8b0f91aca.
//
// Solidity: event ExecuteLimitOrder(uint256 orderId, uint256 dealPrice)
func (_TradingAccount *TradingAccountFilterer) ParseExecuteLimitOrder(log types.Log) (*TradingAccountExecuteLimitOrder, error) {
	event := new(TradingAccountExecuteLimitOrder)
	if err := _TradingAccount.contract.UnpackLog(event, "ExecuteLimitOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingAccountInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TradingAccount contract.
type TradingAccountInitializedIterator struct {
	Event *TradingAccountInitialized // Event containing the contract specifics and raw log

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
func (it *TradingAccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingAccountInitialized)
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
		it.Event = new(TradingAccountInitialized)
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
func (it *TradingAccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingAccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingAccountInitialized represents a Initialized event raised by the TradingAccount contract.
type TradingAccountInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TradingAccount *TradingAccountFilterer) FilterInitialized(opts *bind.FilterOpts) (*TradingAccountInitializedIterator, error) {

	logs, sub, err := _TradingAccount.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TradingAccountInitializedIterator{contract: _TradingAccount.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TradingAccount *TradingAccountFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TradingAccountInitialized) (event.Subscription, error) {

	logs, sub, err := _TradingAccount.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingAccountInitialized)
				if err := _TradingAccount.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TradingAccount *TradingAccountFilterer) ParseInitialized(log types.Log) (*TradingAccountInitialized, error) {
	event := new(TradingAccountInitialized)
	if err := _TradingAccount.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingAccountLimitOpenLongIterator is returned from FilterLimitOpenLong and is used to iterate over the raw logs and unpacked data for LimitOpenLong events raised by the TradingAccount contract.
type TradingAccountLimitOpenLongIterator struct {
	Event *TradingAccountLimitOpenLong // Event containing the contract specifics and raw log

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
func (it *TradingAccountLimitOpenLongIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingAccountLimitOpenLong)
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
		it.Event = new(TradingAccountLimitOpenLong)
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
func (it *TradingAccountLimitOpenLongIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingAccountLimitOpenLongIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingAccountLimitOpenLong represents a LimitOpenLong event raised by the TradingAccount contract.
type TradingAccountLimitOpenLong struct {
	OrderId    *big.Int
	EthSize    *big.Int
	LimitPrice *big.Int
	ExpireAt   *big.Int
	Keeper     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLimitOpenLong is a free log retrieval operation binding the contract event 0x0d3f71d54beaf5776a99a3c0c84a1207de2861ede3d1d85dfcc17df11ebab1d7.
//
// Solidity: event LimitOpenLong(uint256 orderId, uint256 ethSize, uint256 limitPrice, uint256 expireAt, address indexed keeper)
func (_TradingAccount *TradingAccountFilterer) FilterLimitOpenLong(opts *bind.FilterOpts, keeper []common.Address) (*TradingAccountLimitOpenLongIterator, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _TradingAccount.contract.FilterLogs(opts, "LimitOpenLong", keeperRule)
	if err != nil {
		return nil, err
	}
	return &TradingAccountLimitOpenLongIterator{contract: _TradingAccount.contract, event: "LimitOpenLong", logs: logs, sub: sub}, nil
}

// WatchLimitOpenLong is a free log subscription operation binding the contract event 0x0d3f71d54beaf5776a99a3c0c84a1207de2861ede3d1d85dfcc17df11ebab1d7.
//
// Solidity: event LimitOpenLong(uint256 orderId, uint256 ethSize, uint256 limitPrice, uint256 expireAt, address indexed keeper)
func (_TradingAccount *TradingAccountFilterer) WatchLimitOpenLong(opts *bind.WatchOpts, sink chan<- *TradingAccountLimitOpenLong, keeper []common.Address) (event.Subscription, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _TradingAccount.contract.WatchLogs(opts, "LimitOpenLong", keeperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingAccountLimitOpenLong)
				if err := _TradingAccount.contract.UnpackLog(event, "LimitOpenLong", log); err != nil {
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

// ParseLimitOpenLong is a log parse operation binding the contract event 0x0d3f71d54beaf5776a99a3c0c84a1207de2861ede3d1d85dfcc17df11ebab1d7.
//
// Solidity: event LimitOpenLong(uint256 orderId, uint256 ethSize, uint256 limitPrice, uint256 expireAt, address indexed keeper)
func (_TradingAccount *TradingAccountFilterer) ParseLimitOpenLong(log types.Log) (*TradingAccountLimitOpenLong, error) {
	event := new(TradingAccountLimitOpenLong)
	if err := _TradingAccount.contract.UnpackLog(event, "LimitOpenLong", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingAccountLimitOpenShortIterator is returned from FilterLimitOpenShort and is used to iterate over the raw logs and unpacked data for LimitOpenShort events raised by the TradingAccount contract.
type TradingAccountLimitOpenShortIterator struct {
	Event *TradingAccountLimitOpenShort // Event containing the contract specifics and raw log

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
func (it *TradingAccountLimitOpenShortIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingAccountLimitOpenShort)
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
		it.Event = new(TradingAccountLimitOpenShort)
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
func (it *TradingAccountLimitOpenShortIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingAccountLimitOpenShortIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingAccountLimitOpenShort represents a LimitOpenShort event raised by the TradingAccount contract.
type TradingAccountLimitOpenShort struct {
	OrderId    *big.Int
	UsdcSize   *big.Int
	LimitPrice *big.Int
	ExpireAt   *big.Int
	Keeper     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLimitOpenShort is a free log retrieval operation binding the contract event 0xa88905c54863632716f8c0132503b7e5d225a2537cd992c34fef80a5f6c2d800.
//
// Solidity: event LimitOpenShort(uint256 orderId, uint256 usdcSize, uint256 limitPrice, uint256 expireAt, address indexed keeper)
func (_TradingAccount *TradingAccountFilterer) FilterLimitOpenShort(opts *bind.FilterOpts, keeper []common.Address) (*TradingAccountLimitOpenShortIterator, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _TradingAccount.contract.FilterLogs(opts, "LimitOpenShort", keeperRule)
	if err != nil {
		return nil, err
	}
	return &TradingAccountLimitOpenShortIterator{contract: _TradingAccount.contract, event: "LimitOpenShort", logs: logs, sub: sub}, nil
}

// WatchLimitOpenShort is a free log subscription operation binding the contract event 0xa88905c54863632716f8c0132503b7e5d225a2537cd992c34fef80a5f6c2d800.
//
// Solidity: event LimitOpenShort(uint256 orderId, uint256 usdcSize, uint256 limitPrice, uint256 expireAt, address indexed keeper)
func (_TradingAccount *TradingAccountFilterer) WatchLimitOpenShort(opts *bind.WatchOpts, sink chan<- *TradingAccountLimitOpenShort, keeper []common.Address) (event.Subscription, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _TradingAccount.contract.WatchLogs(opts, "LimitOpenShort", keeperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingAccountLimitOpenShort)
				if err := _TradingAccount.contract.UnpackLog(event, "LimitOpenShort", log); err != nil {
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

// ParseLimitOpenShort is a log parse operation binding the contract event 0xa88905c54863632716f8c0132503b7e5d225a2537cd992c34fef80a5f6c2d800.
//
// Solidity: event LimitOpenShort(uint256 orderId, uint256 usdcSize, uint256 limitPrice, uint256 expireAt, address indexed keeper)
func (_TradingAccount *TradingAccountFilterer) ParseLimitOpenShort(log types.Log) (*TradingAccountLimitOpenShort, error) {
	event := new(TradingAccountLimitOpenShort)
	if err := _TradingAccount.contract.UnpackLog(event, "LimitOpenShort", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingAccountNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the TradingAccount contract.
type TradingAccountNewOwnerIterator struct {
	Event *TradingAccountNewOwner // Event containing the contract specifics and raw log

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
func (it *TradingAccountNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingAccountNewOwner)
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
		it.Event = new(TradingAccountNewOwner)
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
func (it *TradingAccountNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingAccountNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingAccountNewOwner represents a NewOwner event raised by the TradingAccount contract.
type TradingAccountNewOwner struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed oldOwner, address indexed newOwner)
func (_TradingAccount *TradingAccountFilterer) FilterNewOwner(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*TradingAccountNewOwnerIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TradingAccount.contract.FilterLogs(opts, "NewOwner", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TradingAccountNewOwnerIterator{contract: _TradingAccount.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed oldOwner, address indexed newOwner)
func (_TradingAccount *TradingAccountFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *TradingAccountNewOwner, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TradingAccount.contract.WatchLogs(opts, "NewOwner", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingAccountNewOwner)
				if err := _TradingAccount.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// ParseNewOwner is a log parse operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed oldOwner, address indexed newOwner)
func (_TradingAccount *TradingAccountFilterer) ParseNewOwner(log types.Log) (*TradingAccountNewOwner, error) {
	event := new(TradingAccountNewOwner)
	if err := _TradingAccount.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingAccountNewPendingOwnerIterator is returned from FilterNewPendingOwner and is used to iterate over the raw logs and unpacked data for NewPendingOwner events raised by the TradingAccount contract.
type TradingAccountNewPendingOwnerIterator struct {
	Event *TradingAccountNewPendingOwner // Event containing the contract specifics and raw log

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
func (it *TradingAccountNewPendingOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingAccountNewPendingOwner)
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
		it.Event = new(TradingAccountNewPendingOwner)
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
func (it *TradingAccountNewPendingOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingAccountNewPendingOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingAccountNewPendingOwner represents a NewPendingOwner event raised by the TradingAccount contract.
type TradingAccountNewPendingOwner struct {
	OldPendingOwner common.Address
	NewPendingOwner common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingOwner is a free log retrieval operation binding the contract event 0xb3d55174552271a4f1aaf36b72f50381e892171636b3fb5447fe00e995e7a37b.
//
// Solidity: event NewPendingOwner(address indexed oldPendingOwner, address indexed newPendingOwner)
func (_TradingAccount *TradingAccountFilterer) FilterNewPendingOwner(opts *bind.FilterOpts, oldPendingOwner []common.Address, newPendingOwner []common.Address) (*TradingAccountNewPendingOwnerIterator, error) {

	var oldPendingOwnerRule []interface{}
	for _, oldPendingOwnerItem := range oldPendingOwner {
		oldPendingOwnerRule = append(oldPendingOwnerRule, oldPendingOwnerItem)
	}
	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _TradingAccount.contract.FilterLogs(opts, "NewPendingOwner", oldPendingOwnerRule, newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TradingAccountNewPendingOwnerIterator{contract: _TradingAccount.contract, event: "NewPendingOwner", logs: logs, sub: sub}, nil
}

// WatchNewPendingOwner is a free log subscription operation binding the contract event 0xb3d55174552271a4f1aaf36b72f50381e892171636b3fb5447fe00e995e7a37b.
//
// Solidity: event NewPendingOwner(address indexed oldPendingOwner, address indexed newPendingOwner)
func (_TradingAccount *TradingAccountFilterer) WatchNewPendingOwner(opts *bind.WatchOpts, sink chan<- *TradingAccountNewPendingOwner, oldPendingOwner []common.Address, newPendingOwner []common.Address) (event.Subscription, error) {

	var oldPendingOwnerRule []interface{}
	for _, oldPendingOwnerItem := range oldPendingOwner {
		oldPendingOwnerRule = append(oldPendingOwnerRule, oldPendingOwnerItem)
	}
	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _TradingAccount.contract.WatchLogs(opts, "NewPendingOwner", oldPendingOwnerRule, newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingAccountNewPendingOwner)
				if err := _TradingAccount.contract.UnpackLog(event, "NewPendingOwner", log); err != nil {
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

// ParseNewPendingOwner is a log parse operation binding the contract event 0xb3d55174552271a4f1aaf36b72f50381e892171636b3fb5447fe00e995e7a37b.
//
// Solidity: event NewPendingOwner(address indexed oldPendingOwner, address indexed newPendingOwner)
func (_TradingAccount *TradingAccountFilterer) ParseNewPendingOwner(log types.Log) (*TradingAccountNewPendingOwner, error) {
	event := new(TradingAccountNewPendingOwner)
	if err := _TradingAccount.contract.UnpackLog(event, "NewPendingOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingAccountOpenLongIterator is returned from FilterOpenLong and is used to iterate over the raw logs and unpacked data for OpenLong events raised by the TradingAccount contract.
type TradingAccountOpenLongIterator struct {
	Event *TradingAccountOpenLong // Event containing the contract specifics and raw log

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
func (it *TradingAccountOpenLongIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingAccountOpenLong)
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
		it.Event = new(TradingAccountOpenLong)
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
func (it *TradingAccountOpenLongIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingAccountOpenLongIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingAccountOpenLong represents a OpenLong event raised by the TradingAccount contract.
type TradingAccountOpenLong struct {
	EthSize      *big.Int
	CTokenMint   *big.Int
	BorrowAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOpenLong is a free log retrieval operation binding the contract event 0xefb1650edbd2ee9b7b1940e049e74fc26ce6f3d0f62f1a3f32e25397b4dffe3b.
//
// Solidity: event OpenLong(uint256 ethSize, uint256 cTokenMint, uint256 borrowAmount)
func (_TradingAccount *TradingAccountFilterer) FilterOpenLong(opts *bind.FilterOpts) (*TradingAccountOpenLongIterator, error) {

	logs, sub, err := _TradingAccount.contract.FilterLogs(opts, "OpenLong")
	if err != nil {
		return nil, err
	}
	return &TradingAccountOpenLongIterator{contract: _TradingAccount.contract, event: "OpenLong", logs: logs, sub: sub}, nil
}

// WatchOpenLong is a free log subscription operation binding the contract event 0xefb1650edbd2ee9b7b1940e049e74fc26ce6f3d0f62f1a3f32e25397b4dffe3b.
//
// Solidity: event OpenLong(uint256 ethSize, uint256 cTokenMint, uint256 borrowAmount)
func (_TradingAccount *TradingAccountFilterer) WatchOpenLong(opts *bind.WatchOpts, sink chan<- *TradingAccountOpenLong) (event.Subscription, error) {

	logs, sub, err := _TradingAccount.contract.WatchLogs(opts, "OpenLong")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingAccountOpenLong)
				if err := _TradingAccount.contract.UnpackLog(event, "OpenLong", log); err != nil {
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

// ParseOpenLong is a log parse operation binding the contract event 0xefb1650edbd2ee9b7b1940e049e74fc26ce6f3d0f62f1a3f32e25397b4dffe3b.
//
// Solidity: event OpenLong(uint256 ethSize, uint256 cTokenMint, uint256 borrowAmount)
func (_TradingAccount *TradingAccountFilterer) ParseOpenLong(log types.Log) (*TradingAccountOpenLong, error) {
	event := new(TradingAccountOpenLong)
	if err := _TradingAccount.contract.UnpackLog(event, "OpenLong", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingAccountOpenShortIterator is returned from FilterOpenShort and is used to iterate over the raw logs and unpacked data for OpenShort events raised by the TradingAccount contract.
type TradingAccountOpenShortIterator struct {
	Event *TradingAccountOpenShort // Event containing the contract specifics and raw log

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
func (it *TradingAccountOpenShortIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingAccountOpenShort)
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
		it.Event = new(TradingAccountOpenShort)
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
func (it *TradingAccountOpenShortIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingAccountOpenShortIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingAccountOpenShort represents a OpenShort event raised by the TradingAccount contract.
type TradingAccountOpenShort struct {
	UsdcSize     *big.Int
	CTokenMint   *big.Int
	BorrowAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOpenShort is a free log retrieval operation binding the contract event 0x383a754f7fb836d68aa84a08d21e9abdd63725207f15a8d8777e1992ca92f4ed.
//
// Solidity: event OpenShort(uint256 usdcSize, uint256 cTokenMint, uint256 borrowAmount)
func (_TradingAccount *TradingAccountFilterer) FilterOpenShort(opts *bind.FilterOpts) (*TradingAccountOpenShortIterator, error) {

	logs, sub, err := _TradingAccount.contract.FilterLogs(opts, "OpenShort")
	if err != nil {
		return nil, err
	}
	return &TradingAccountOpenShortIterator{contract: _TradingAccount.contract, event: "OpenShort", logs: logs, sub: sub}, nil
}

// WatchOpenShort is a free log subscription operation binding the contract event 0x383a754f7fb836d68aa84a08d21e9abdd63725207f15a8d8777e1992ca92f4ed.
//
// Solidity: event OpenShort(uint256 usdcSize, uint256 cTokenMint, uint256 borrowAmount)
func (_TradingAccount *TradingAccountFilterer) WatchOpenShort(opts *bind.WatchOpts, sink chan<- *TradingAccountOpenShort) (event.Subscription, error) {

	logs, sub, err := _TradingAccount.contract.WatchLogs(opts, "OpenShort")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingAccountOpenShort)
				if err := _TradingAccount.contract.UnpackLog(event, "OpenShort", log); err != nil {
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

// ParseOpenShort is a log parse operation binding the contract event 0x383a754f7fb836d68aa84a08d21e9abdd63725207f15a8d8777e1992ca92f4ed.
//
// Solidity: event OpenShort(uint256 usdcSize, uint256 cTokenMint, uint256 borrowAmount)
func (_TradingAccount *TradingAccountFilterer) ParseOpenShort(log types.Log) (*TradingAccountOpenShort, error) {
	event := new(TradingAccountOpenShort)
	if err := _TradingAccount.contract.UnpackLog(event, "OpenShort", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingAccountWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the TradingAccount contract.
type TradingAccountWithdrawIterator struct {
	Event *TradingAccountWithdraw // Event containing the contract specifics and raw log

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
func (it *TradingAccountWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingAccountWithdraw)
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
		it.Event = new(TradingAccountWithdraw)
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
func (it *TradingAccountWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingAccountWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingAccountWithdraw represents a Withdraw event raised by the TradingAccount contract.
type TradingAccountWithdraw struct {
	IsETH            bool
	CTokenAmount     *big.Int
	UnderlyingAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfedb0622d5d5fa46cc706a5c2c169f2bf2eb31a59ff2f70cef3096376f58c2c0.
//
// Solidity: event Withdraw(bool isETH, uint256 cTokenAmount, uint256 underlyingAmount)
func (_TradingAccount *TradingAccountFilterer) FilterWithdraw(opts *bind.FilterOpts) (*TradingAccountWithdrawIterator, error) {

	logs, sub, err := _TradingAccount.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &TradingAccountWithdrawIterator{contract: _TradingAccount.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfedb0622d5d5fa46cc706a5c2c169f2bf2eb31a59ff2f70cef3096376f58c2c0.
//
// Solidity: event Withdraw(bool isETH, uint256 cTokenAmount, uint256 underlyingAmount)
func (_TradingAccount *TradingAccountFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *TradingAccountWithdraw) (event.Subscription, error) {

	logs, sub, err := _TradingAccount.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingAccountWithdraw)
				if err := _TradingAccount.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfedb0622d5d5fa46cc706a5c2c169f2bf2eb31a59ff2f70cef3096376f58c2c0.
//
// Solidity: event Withdraw(bool isETH, uint256 cTokenAmount, uint256 underlyingAmount)
func (_TradingAccount *TradingAccountFilterer) ParseWithdraw(log types.Log) (*TradingAccountWithdraw, error) {
	event := new(TradingAccountWithdraw)
	if err := _TradingAccount.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
