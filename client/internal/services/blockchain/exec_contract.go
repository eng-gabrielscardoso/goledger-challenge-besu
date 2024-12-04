package blockchain

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ExecContract(method string, value uint64) {
	abi, err := abi.JSON(strings.NewReader(os.Getenv("SIMPLE_STORAGE_ABI")))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := ethclient.DialContext(ctx, os.Getenv("SIMPLE_STORAGE_NETWORK_URL"))

	if err != nil {
		log.Fatalf("error dialing node: %v", err)
	}

	slog.Info("querying chain id")

	chainId, err := client.ChainID(ctx)

	if err != nil {
		log.Fatalf("error querying chain id: %v", err)
	}

	defer client.Close()

	contractAddress := common.HexToAddress(os.Getenv("SIMPLE_STORAGE_CONTRACT_ADDRESS"))

	boundContract := bind.NewBoundContract(
		contractAddress,
		abi,
		client,
		client,
		client,
	)

	priv, err := crypto.HexToECDSA(os.Getenv("SIMPLE_STORAGE_PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("error loading private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(priv, chainId)
	if err != nil {
		log.Fatalf("error creating transactor: %v", err)
	}

	tx, err := boundContract.Transact(auth, method, big.NewInt((int64(value))))
	if err != nil {
		log.Fatalf("error transacting: %v", err)
	}

	fmt.Println("waiting until transaction is mined",
		"tx", tx.Hash().Hex(),
	)

	receipt, err := bind.WaitMined(
		context.Background(),
		client,
		tx,
	)

	if err != nil {
		log.Fatalf("error waiting for transaction to be mined: %v", err)
	}

	fmt.Printf("transaction mined: %v\n", receipt)
}
