package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
	"net/http"
	"strconv"
)

const (
	EthereumNetwork = "https://ropsten.infura.io"
	EvvTestedWallet = "0x66623A091684C70d3B6fdc5a1222C448B5b3B365"
	NtnTestedWallet = "0xc903EB80d685091Da87ab2ffD10A594A0EAd6522"
	CLIENT          = "0xbE34A96D650F318d85c6584d376930012A6d1F78"
)

var client *ethclient.Client

func EtherPerUsd() float64 {

	type eth_to_usd struct {
		Usd string `json:"price_usd"`
	}

	const ether_link = "https://api.coinmarketcap.com/v1/ticker/ethereum/"
	const usd_link = "https://free.currconv.com/api/v7/convert?q=USD_RUB&compact=ultra&apiKey=60350c5cceac35a4ca7e"
	resp, err := http.Get(ether_link)
	if err != nil {
		fmt.Println("We have some error in getting exchange rate ETH / USD ", err)
		return -1
	}
	var rate []eth_to_usd
	err = json.NewDecoder(resp.Body).Decode(&rate)
	if err != nil {
		fmt.Println("Decode error: ", err)
		return -1
	}
	usd_rate, err := strconv.ParseFloat(rate[0].Usd, 64)
	return usd_rate
}

func CreateWallet() string {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 9a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address) // 0x96216849c49358B10257cb55b28eA603c874b05E

	return address
}

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
	balance, avalBalance, err := CheckBalance(EvvTestedWallet)
	fmt.Println(balance, avalBalance, err)

	//client := CreateWallet()
	balance, avalBalance, err = CheckBalance(CLIENT)
	fmt.Println(balance, avalBalance, err)

	fmt.Println(EtherPerUsd())

	//client, err := ethclient.Dial("https://ropsten.infura.io")

	//pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	//fmt.Println(pendingBalance) // 25729324269165216042
}
