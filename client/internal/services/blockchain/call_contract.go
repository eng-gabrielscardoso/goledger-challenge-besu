package blockchain

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CallContract() {
	var result interface{}

	abi, err := abi.JSON(strings.NewReader("REPLACE: abi JSON as string goes here"))

	if err != nil {
		log.Fatalf("error parsing abi: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := ethclient.DialContext(ctx, "REPLACE: network URL")

	if err != nil {
		log.Fatalf("error connecting to eth client: %v", err)
	}

	defer client.Close()

	contractAddress := common.HexToAddress("REPLACE: contract address")

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

	err = boundContract.Call(&caller, &output, "REPLACE: method name")

	if err != nil {
		log.Fatalf("error calling contract: %v", err)
	}

	result = output

	fmt.Println("Successfully called contract!", result)
}
