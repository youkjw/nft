package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
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

	//查看余额
	formAddress := common.HexToAddress("0x721ceEFB1F0B5DDBd7450306C44b06841489B375")
	balance, err := instance.BalanceOf(&bind.CallOpts{}, formAddress)
	fmt.Printf("wei: %s\n", balance.String()) // "wei"

	privateKey, err := crypto.HexToECDSA("4d258c34b98d3c1fb0d1e70201cdee95ce37150179d34e5ec0083b3143b93136")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("gasPrice:%s", gasPrice.String())

	toAddress := common.HexToAddress("0x14857cAFF4be7B441b525723813e64655f67f366")
	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10) // 1000 tokens
	data, err := MakeERC20TransferData(toAddress, amount)
	if err != nil {
		log.Fatal(err)
	}

	gasLimit := uint64(30000000)
	//gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	//	To:   &toAddress,
	//	Data: data,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Printf("gasLimit:%d", gasLimit)

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &tokenAddress,
		Value:    big.NewInt(0), //eth value
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	})

	//chainID, err := client.NetworkID(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	chainID := big.NewInt(1337)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

func SignTransaction(tx *types.Transaction, privateKeyStr string, chainId *big.Int) (string, error) {
	privateKey, err := StringToPrivateKey(privateKeyStr)
	if err != nil {
		return "", err
	}
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	//signTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		return "", nil
	}

	b, err := rlp.EncodeToBytes(signTx)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func StringToPrivateKey(privateKeyStr string) (*ecdsa.PrivateKey, error) {
	privateKeyByte, err := hexutil.Decode(privateKeyStr)
	if err != nil {
		return nil, err
	}
	privateKey, err := crypto.ToECDSA(privateKeyByte)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func KeystoreToPrivateKey(keystoreContent []byte, password string) (*ecdsa.PrivateKey, error) {
	unlockedKey, err := keystore.DecryptKey(keystoreContent, password)
	if err != nil {
		return nil, err
	}
	return unlockedKey.PrivateKey, nil
}

func MakeERC20TransferData(toAddress common.Address, amount *big.Int) ([]byte, error) {
	var data []byte

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodId := hash.Sum(nil)[:4]

	data = append(data, methodId...)
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	data = append(data, paddedAddress...)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	data = append(data, paddedAmount...)
	return data, nil
}

func OfflineTransferERC20(chainID *big.Int, nonce uint64, toAddress common.Address, toContractAddress string, value *big.Int, privk string, gasLimit uint64, gasPrice *big.Int) (string, error) {
	data, err := MakeERC20TransferData(toAddress, value)
	if err != nil {
		return "", err
	}

	tx := types.NewTransaction(uint64(nonce), common.HexToAddress(toContractAddress), big.NewInt(0), gasLimit, gasPrice, data)
	return SignTransaction(tx, privk, chainID)
}
