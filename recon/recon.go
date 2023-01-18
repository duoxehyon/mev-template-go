package recon

import (
	"context"
	"fmt"
	"mev-template-go/types"

	"github.com/ethereum/go-ethereum/common"
	geth_types "github.com/ethereum/go-ethereum/core/types"
)

func AlertTransaction(cfg types.Config, WatchTo map[common.Address]bool, sendTo chan *geth_types.Transaction) {

	pendingTransactions := make(chan common.Hash)
	sub, err := cfg.RpcClient.EthSubscribe(context.Background(), pendingTransactions, "newPendingTransactions")
	if err != nil {
		return
	}

	for {
		txHash := <-pendingTransactions

		txnBody, isPending, err := cfg.Client.TransactionByHash(context.Background(), txHash)
		if err != nil || !isPending {
			continue
		}

		if txnBody.To() == nil {
			continue
		}

		val := WatchTo[*txnBody.To()]
		if !val {
			continue
		}

		sendTo <- txnBody
	}

	defer sub.Unsubscribe()
}

func AlertBlocks(cfg types.Config) {

	headers := make(chan *geth_types.Header)
	sub, err := cfg.ClientWss.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		return
	}
	for {
		select {
		case header := <-headers:
			fmt.Println("------- block: ", header.Number.Uint64(), " -------")
		}

	}

	defer sub.Unsubscribe()
}
