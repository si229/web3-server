package handlers

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/si229/web3-api/blockchain"
	"github.com/si229/web3-api/internal/config"
	request "github.com/si229/web3-api/internal/model"
	"github.com/si229/web3-api/internal/utils"
	"github.com/si229/web3-api/internal/utils/logger"
)

func Deposit(c *gin.Context) {

	var req request.DepositReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	logger.Log.Info().Msgf("body:%+v", req)

	chainID := big.NewInt(31337)

	auth, err := blockchain.NewAuth(req.PrivateKey, chainID)

	tokenAmount := utils.ToTokenUnits(req.Amount, config.TokenDecimals[req.Symbol])

	if err != nil {
		logger.Log.Error().Msgf("deposit  error %+v", err)
		c.JSON(200, gin.H{"data": "fail"})
	}

	if req.Symbol == "ETH" {
		auth.Value = tokenAmount
		_, err := blockchain.GetBaccarat().HTTP.Deposit(auth)
		if err != nil {
			logger.Log.Info().Msgf("deposit  error %+v", err)
			c.JSON(400, gin.H{"error": err.Error()})
		} else {
			logger.Log.Info().Msgf("deposit %f", utils.ToEth(auth.Value))
			c.JSON(200, gin.H{"data": "ok"})
		}
	} else {
		_, err := blockchain.GetBaccarat().HTTP.Deposit0(auth, config.SymbolToken[req.Symbol], tokenAmount)
		if err != nil {
			c.JSON(200, gin.H{"data": "fail"})
		} else {
			logger.Log.Info().Msgf("deposit %f", utils.ToEth(auth.Value))
			c.JSON(200, gin.H{"data": "ok"})
		}

	}

}

func Withdraw(c *gin.Context) {
	var req request.WithdrawReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	logger.Log.Info().Msgf("body:%+v", req)

	chainID := big.NewInt(31337)

	auth, err := blockchain.NewAuth(req.PrivateKey, chainID)

	token := config.SymbolToken[req.Symbol]
	tokenAmount := utils.ToTokenUnits(req.Amount, config.TokenDecimals[req.Symbol])
	tx, err := blockchain.GetBaccarat().HTTP.Withdraw0(auth, token, tokenAmount)
	if err != nil {
		logger.Log.Info().Msgf("withdraw error:%s+v", err)
		c.JSON(200, gin.H{"error": err.Error()})
		return
	}
	logger.Log.Info().Msgf("balance Withdraw:%s", tx.Hash().Hex())
	c.JSON(200, gin.H{"result": "ok"})
}

func GetBalance(c *gin.Context) {
	var req request.GetBalanceReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	logger.Log.Info().Msgf("body:%+v", req)

	token := config.SymbolToken[req.Symbol]

	bal, err := blockchain.GetBaccarat().HTTP.GetBalance0(
		&bind.CallOpts{From: common.HexToAddress(req.Address)}, token)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	balance := utils.NormalizeTokenAmount(bal, config.TokenDecimals[req.Symbol])
	c.JSON(200, gin.H{"data": balance})
}

func GetContractBalance(c *gin.Context) {

	var req request.GetContractBalanceReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	logger.Log.Info().Msgf("body:%+v", req)
	// 2. 合约地址
	contractAddress := common.HexToAddress(req.Address)
	// 3. 获取余额
	balance, err := blockchain.GetHTTPClient().BalanceAt(context.Background(), contractAddress, nil)

	if err != nil {
		logger.Log.Error().Msgf("body:%+v", err)
	}
	bal := utils.NormalizeTokenAmount(balance, config.TokenDecimals["ETH"])
	c.JSON(200, gin.H{"data": bal})
}
