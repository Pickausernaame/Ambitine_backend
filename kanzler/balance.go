package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

const (
	EthereumNetwork = "https://ropsten.infura.io"
	EvvTestedWallet = "0x66623A091684C70d3B6fdc5a1222C448B5b3B365"
	NtnTestedWallet = "0xc903EB80d685091Da87ab2ffD10A594A0EAd6522"
)

var client *ethclient.Client

func BlockchainClientInit(netAddress string) (connectionErr error) {
	client, connectionErr = ethclient.Dial(netAddress)
	return
}

func BigIntConvert(value *big.Int) (convertedValue *big.Float) {
	fbalance := new(big.Float)
	fbalance.SetString(value.String())
	convertedValue = new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return
}

func CheckBalance(walletAddress string) (*big.Float, *big.Float, error) {
	account := common.HexToAddress(walletAddress)
	currentIntBalance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, nil, err
	}
	pendingIntBalance, err := client.PendingBalanceAt(context.Background(), account)

	return BigIntConvert(pendingIntBalance), BigIntConvert(currentIntBalance), err
}

func main() {
	err := BlockchainClientInit(EthereumNetwork)
	if err != nil {
		log.Fatal(err)
	}
	balance, avalBalance, err := CheckBalance(NtnTestedWallet)
	fmt.Println(balance, avalBalance, err)
	//client, err := ethclient.Dial("https://ropsten.infura.io")

	//pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	//fmt.Println(pendingBalance) // 25729324269165216042
}
