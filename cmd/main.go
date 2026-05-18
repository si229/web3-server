package main

import (
	"github.com/gin-gonic/gin"
	"github.com/si229/web3-api/blockchain"
	"github.com/si229/web3-api/internal/config"
	router "github.com/si229/web3-api/internal/routers"
	"github.com/si229/web3-api/internal/utils/logger"
	"github.com/si229/web3-api/internal/web3"
	"github.com/si229/web3-api/internal/worker"
)

var AccountPrivateKey = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
var ContractAddress = "0x8464135c8F25Da09e49BC8782676a84730C318bC"
var ContractUSDTAddress = "0x71C95911E9a5D330f4D621842EC243EE1343292e"

func main() {
	logger.Init()
	web3.Init()

	blockchain.InitWSClient("ws://127.0.0.1:8545")
	blockchain.InitHTTPClient("http://127.0.0.1:8545")
	blockchain.InitBaccarat(ContractAddress)
	go worker.StartBaccaratListener()
	startHttp()
	select {}
}

func startHttp() {
	r := gin.Default()
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	router.RegisterRoutes(r, cfg)
	r.Run(":9090")
}
