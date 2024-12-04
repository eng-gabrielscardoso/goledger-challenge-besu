package blockchain

import (
	"context"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/eng-gabrielscardoso/goledger-challenge-besu/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CallContract(method string) *big.Int {
	var result *big.Int

	abi, err := abi.JSON(strings.NewReader(os.Getenv("SIMPLE_STORAGE_ABI")))

	if err != nil {
		log.Fatalf("Error parsing abi: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := ethclient.DialContext(ctx, os.Getenv("SIMPLE_STORAGE_NETWORK_URL"))

	if err != nil {
		log.Fatalf("Error connecting to eth client: %v", err)
	}

	defer client.Close()

	contractAddress := common.HexToAddress(os.Getenv("SIMPLE_STORAGE_CONTRACT_ADDRESS"))

	caller := bind.CallOpts{
		Pending: false,
		Context: ctx,
	}

	boundContract := bind.NewBoundContract(
		contractAddress,
		abi,
		client,
		client,
		client,
	)

	var output []interface{}

	err = boundContract.Call(&caller, &output, method)

	if err != nil {
		log.Fatalf("Error calling contract: %v", err)
	}

	if len(output) > 0 {
		if val, ok := output[0].(*big.Int); ok {
			result = val
		} else {
			log.Fatalf("Unexpected output type: %v", output[0])
		}
	} else {
		log.Fatalf("No output returned from contract call")
	}

	utils.Logger.Printf("Successfully called contract with result: %v", result)

	return result
}
