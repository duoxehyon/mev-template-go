package main

import (
	"crypto/ecdsa"
	"fmt"
	"os"

	"mev-template-go/contract_modules/uniswap_v2"
	"mev-template-go/recon"
	"mev-template-go/types"

	"context"

	"github.com/ethereum/go-ethereum/common"
	geth_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	log "github.com/inconshreveable/log15"
	"github.com/joho/godotenv"
)

var config types.Config

// Initialize function to initialize the client, private key, and wallet address
func Initialize() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	walletAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	client, err := ethclient.Dial(os.Getenv("NETWORK_RPC"))
	if err != nil {
		return err
	}

	clientWss, err := ethclient.Dial(os.Getenv("NETWORK_WSS"))
	if err != nil {
		return err
	}

	rpcClient, err := rpc.DialContext(context.Background(), os.Getenv("NETWORK_WSS"))
	if err != nil {
		return err
	}

	config = types.Config{
		Client:        *client,
		ClientWss:     *clientWss,
		RpcClient:     *rpcClient,
		PrivateKey:    privateKey,
		WalletAddress: walletAddress,
	}
	return nil
}

func main() {

	// Initialize the client, private key, and wallet address
	err := Initialize()
	if err != nil {
		log.Error("Initialization failed", "error", err)
	}

	fmt.Println("MEV TEMPLATE IN GO")
	fmt.Println("")

	uniV2, err := uniswap_v2.New(config, common.HexToAddress(uniV2RouterAddress), common.HexToAddress(uniV2FactoryAddress), "contracts/bindings/uniswap_v2/uniswap_v2_router/UniV2Router.json")

	receiveChannel := make(chan *geth_types.Transaction)
	go recon.AlertTransaction(config, map[common.Address]bool{common.HexToAddress(uniV2RouterAddress): true}, receiveChannel)
	go recon.AlertBlocks(config)

	for {
		newTx := <-receiveChannel
		uniV2.DecodeTransactionInputData(newTx.Data())
	}

}
