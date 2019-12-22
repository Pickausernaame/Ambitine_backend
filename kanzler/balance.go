package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

const (
	EthereumNetwork        = "https://ropsten.infura.io"
	ETH_EXCHANGE_RATE_LINK = "https://api.coinmarketcap.com/v1/ticker/ethereum/"

	EvvTestedWallet = "0x66623A091684C70d3B6fdc5a1222C448B5b3B365"
	NtnTestedWallet = "0xc903EB80d685091Da87ab2ffD10A594A0EAd6522"
	PRIVATE_KEY_NTN = "e9ef28ce8d86c134564c3b9c0ea2d4180d59a645e186cceb84fbfb7b6638d07b"

	CREATED_PRIVATE_KEY = "8c291af4304e84fb5f7221ed5c3e147566b4b94022fc38c5c33e374be8683944"
	CREATED_ADRESS      = "0xe5b7E4070a7Ebd127eEff3a8C1c71f4ae55d516A"

	ETH       = 1000000000000000000
	GAS_LIMIT = 21000
)

var client *ethclient.Client

func private_to_public(privateKey *ecdsa.PrivateKey) (fromAddress common.Address, err error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
	}
	fromAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
	return
}

func MakeTransaction(fromPrivateKey string, toAddress string, amount float64) (err error) {
	privateKey, err := crypto.HexToECDSA(fromPrivateKey)
	if err != nil {
		return
	}

	fromAddress, err := private_to_public(privateKey)
	if err != nil {
		return
	}
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return
	}
	final_amount := int64(amount * ETH)
	value := big.NewInt(final_amount)
	gasLimit := uint64(GAS_LIMIT)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	var data []byte
	tx := types.NewTransaction(nonce, common.HexToAddress(toAddress), value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return
	}
	return nil
}

func EtherPerUsd() float64 {

	type eth_to_usd struct {
		Usd string `json:"price_usd"`
	}

	resp, err := http.Get(ETH_EXCHANGE_RATE_LINK)
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

func CreateWallet() (privateKey string, address string, err error) {
	privateKeyECDSA, err := crypto.GenerateKey()
	if err != nil {
		return
	}

	privateKeyBytes := crypto.FromECDSA(privateKeyECDSA)
	privateKey = hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
	}

	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return privateKey, address, nil
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
	_, avalBalance, err := CheckBalance(EvvTestedWallet)
	//fmt.Println(balance, avalBalance, err)

	//privateKey, adress := CreateWallet()
	//fmt.Println(privateKey)
	//fmt.Println(adress)

	//balance, avalBalance, err = CheckBalance(CLIENT)
	//fmt.Println(balance, avalBalance, err)

	MakeTransaction(PRIVATE_KEY_NTN, CREATED_ADRESS, 0.1)
	time.Sleep(60 * time.Second)
	_, avalBalance, err = CheckBalance(CREATED_ADRESS)
	fmt.Println("BALANCE OF CREATED WALLET AFTER TRANSACTION OF 0.1 ETH FROM NTN", avalBalance)
	_, avalBalance, err = CheckBalance(NtnTestedWallet)
	fmt.Println("BALANCE OF NTN_WALLET AFTER TRANSACTION OF 0.1 ETH TO CREATED_WALLET", avalBalance)

	MakeTransaction(CREATED_PRIVATE_KEY, NtnTestedWallet, 0.1)
	time.Sleep(90 * time.Second)
	_, avalBalance, err = CheckBalance(CREATED_ADRESS)
	fmt.Println("BALANCE OF CREATED WALLET AFTER TRANSACTION OF 0.1 ETH TO NTN", avalBalance)
	_, avalBalance, err = CheckBalance(NtnTestedWallet)
	fmt.Println("BALANCE OF NTN_WALLET AFTER TRANSACTION OF 0.1 ETH FROM CREATED_WALLET", avalBalance)

	//client, err := ethclient.Dial("https://ropsten.infura.io")

	//pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	//fmt.Println(pendingBalance) // 25729324269165216042
}
