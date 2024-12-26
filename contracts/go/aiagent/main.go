package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	ctx := context.Background()

	client, err := ethclient.Dial("ws://localhost:8546")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x676E400d0200Ac8f3903A3CDC7cc3feaF21004d0")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Subscribed to contract logs for contract address: ", contractAddress.Hex())

	for {
		select {
		case err = <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			log.Println(vLog) // pointer to event log
		}
	}
}
