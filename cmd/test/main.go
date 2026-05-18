package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/si229/web3-api/blockchain"
	"github.com/si229/web3-api/internal/config"
	"github.com/si229/web3-api/internal/utils"
)

var ContractAddress = "0x8464135c8F25Da09e49BC8782676a84730C318bC"

func main() {

	blockchain.InitHTTPClient("http://127.0.0.1:8545")
	blockchain.InitBaccarat(ContractAddress)
	// 私钥
	privateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		log.Fatal(err)
	}

	// 公钥地址
	fromAddress := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	// 获取 nonce
	nonce, err := blockchain.GetHTTPClient().PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 设置 gas 参数
	gasLimit := uint64(300000) // 根据合约函数复杂度调整
	gasPrice, err := blockchain.GetHTTPClient().SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 合约地址
	contractAddress := common.HexToAddress(ContractAddress)

	// 已经生成好的 data
	data := common.FromHex("0xf4d4c9d700000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001bc16d674ec80000")

	amount := utils.ToTokenUnits(2, config.TokenDecimals["ETH"])

	// 构造交易
	tx := types.NewTransaction(
		nonce,
		contractAddress,
		amount,
		gasLimit,
		gasPrice,
		data,
	)

	// 签名交易
	chainID, err := blockchain.GetHTTPClient().NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 发送交易
	err = blockchain.GetHTTPClient().SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("交易已发送: %s\n", signedTx.Hash().Hex())
}
