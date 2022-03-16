package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"nft"
)

func main() {
	client, err := ethclient.Dial("http://175.178.183.199:8545")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x2848F4E83c234EfDeC053A39E3cCB313401E0767")
	instance, err := nft.NewNft(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	_ = instance
}


