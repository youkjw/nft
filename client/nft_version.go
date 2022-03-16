package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"nft"
)

func main() {
	ctx := context.Background()
	client, err := ethclient.Dial("http://175.178.183.199:8545")
	if err != nil {
		log.Fatal(err)
	}

	blockNum, err := client.BlockNumber(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(blockNum)

	block, err := client.BlockByNumber(ctx, big.NewInt(int64(blockNum)))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(block)

	address := common.HexToAddress("0xa3F28bBcd5C4740a9dEaCE8909b40aFcDfD40373")
	instance, err := nft.NewNft(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")

	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version) // "1.0"
}
