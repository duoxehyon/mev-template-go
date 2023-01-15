package types

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type Config struct {
	Client        ethclient.Client
	PrivateKey    *ecdsa.PrivateKey
	WalletAddress common.Address
}

type UniV2Pool struct {
	Address common.Address

	Token0 common.Address
	Token1 common.Address

	Reserve0 big.Int
	Reserve1 big.Int

	Fees uint16
}
