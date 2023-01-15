package main

import (
	"crypto/ecdsa"
	"fmt"
	"os"

	"mev-template-go/contract_modules/uniswap_v2"
	"mev-template-go/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/inconshreveable/log15"
	"github.com/joho/godotenv"
)

var config types.Config

// Initialize function to initialize the client, private key, and wallet address
func Initialize() error {
	// Load the env file
	if err := godotenv.Load(".env"); err != nil {
		return err
	}

	// Get the private key from env
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return err
	}

	// Get the public key from private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// Get the wallet address from public key
	walletAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Connect to the RPC
	client, err := ethclient.Dial(os.Getenv("NETWORK_RPC"))
	if err != nil {
		return err
	}

	config = types.Config{
		Client:        *client,
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

	stuff, err := uniswap_v2.New(config, common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"), common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"), "contracts/bindings/uniswap_v2/uniswap_v2_router/UniV2Router.json")

}
