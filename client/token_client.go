package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"nft/token"
)

func main() {
	client, err := ethclient.Dial("http://175.178.183.199:8545")
	if err != nil {
		log.Fatal(err)
	}

	tokenAddress := common.HexToAddress("0x2d9F24Ed5BD0a6d93662936Fcd00c456991F5483")
	instance, err := token.NewCm(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	address := common.HexToAddress("0x721ceEFB1F0B5DDBd7450306C44b06841489B375")
	balance, err := instance.BalanceOf(&bind.CallOpts{}, address)
	fmt.Printf("wei: %s\n", balance.String()) // "wei"
}
