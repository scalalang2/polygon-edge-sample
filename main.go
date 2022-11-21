package main

import (
	"context"
	"math/big"
	"os"
	"polygon-edge-sample/gen"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/inconshreveable/log15"
)

var (
	privateKey  = "2feea02f07cdc6d3e9528961615821e5b24e8f6c26fb866939c2743d74d9a8e2"
	userAddress = "0xb507bae31ba521ad00a718d228045aba2d8832e6"
	alice       = "0x5a9E4bE2A323b95883fDe23F8404891C4068321a"
	chainId     = int64(100)
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Error("failed to connect to the polygon edge node", "err", err)
	}

	opt, _ := buildOpt(privateKey, chainId)
	contractAddress, tx, token, err := gen.DeployToken(opt, client)
	if err != nil {
		log.Error("failed to deploy token contract", "err", err)
		os.Exit(1)
	}

	waitUntilTxCommitted(client, tx.Hash())
	log.Info("deployed token contract to polygon edge", "contractAddr", contractAddress)

	// transfer 100 token to alice
	balance, err := token.BalanceOf(&bind.CallOpts{}, common.HexToAddress(userAddress))
	log.Info("[before] balanceOf alice", "balance", balance)

	tx, err = token.Transfer(opt, common.HexToAddress(alice), big.NewInt(100))
	if err != nil {
		log.Error("failed to execute 'transfer' method", "err", err)
		os.Exit(1)
	}

	waitUntilTxCommitted(client, tx.Hash())
	balance, err = token.BalanceOf(&bind.CallOpts{}, common.HexToAddress(alice))

	log.Info("[after] balanceOf alice", "balance", balance)
}

func buildOpt(privateKey string, chainId int64) (*bind.TransactOpts, error) {
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(pk, big.NewInt(chainId))
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func waitUntilTxCommitted(client *ethclient.Client, hash common.Hash) {
	tx, pending, err := client.TransactionByHash(context.Background(), hash)
	if err == nil && !pending {
		log.Info("tx committed", "txhash", tx.Hash().Hex())
		return
	}

	if err != nil {
		log.Error("unable to get transaction", "err", err)
	}

	waitUntilTxCommitted(client, hash)
}
