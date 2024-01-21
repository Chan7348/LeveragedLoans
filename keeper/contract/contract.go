package contract

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"

	"limit/conf"
)

var (
	client        *ethclient.Client
	auth          *bind.TransactOpts
	walletAddress common.Address
	priceFeed     *PriceFeed
)

func Init() {
	log.Println("Init contract ....")
	c, err := ethclient.Dial(conf.Config.RPC)
	if err != nil {
		panic(err)
	}
	client = c

	privateKey, err := crypto.HexToECDSA(conf.Config.PrivateKey)
	if err != nil {
		log.Print(err)
	}

	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(conf.Config.ChainID)))
	if err != nil {
		log.Print(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Print("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	walletAddress = crypto.PubkeyToAddress(*publicKeyECDSA)

	priceFeed = newPriceFeed()
}

func getNonce() uint64 {
	nonce, err := client.PendingNonceAt(context.Background(), walletAddress)
	if err != nil {
		log.Printf("get nonce error: %s", err)
		return 0
	}
	return nonce
}

func getGasPrice() *big.Int {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Printf("get gas price error: %s", err)
		return common.Big0
	}
	return gasPrice
}

// 实例化 PriceOracle 合约
func newPriceFeed() *PriceFeed {
	priceFeedAddress := common.HexToAddress(conf.Config.PriceFeed)
	instance, err := NewPriceFeed(priceFeedAddress, client)
	if err != nil {
		log.Print(err)
	}
	return instance
}

func GetWalletAddress() string {
	return walletAddress.Hex()
}

// 获取 ETH/UNI 的价格
func GetLatestPrice() decimal.Decimal {
	result, _ := priceFeed.LatestRoundData(nil)
	return decimal.NewFromBigInt(result.Answer, 8)
}

func ExecuteLimitOrder(tradingAccount string, orderId uint) (string, error) {
	accountContract, err := NewTradingAccount(common.HexToAddress(tradingAccount), client)
	if err != nil {
		log.Printf("NewUEther error: %s", err)
		return "", err
	}

	auth.Nonce = big.NewInt(int64(getNonce()))
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = getGasPrice()

	tx, err := accountContract.ExecuteLimitOrder(auth, big.NewInt(int64(orderId)))
	if err != nil {
		log.Printf("LiquidateBorrowETH error: %s", err)
		return "", err
	}

	return tx.Hash().String(), nil
}
