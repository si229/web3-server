package service

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/si229/web3-api/internal/client"
	"github.com/si229/web3-api/internal/config"
	contract "github.com/si229/web3-api/internal/contract"
	"github.com/si229/web3-api/internal/utils"
	"github.com/si229/web3-api/internal/utils/logger"
)

func HandleUpdateBalance(player common.Address, token uint8, instance *contract.Baccarat) {
	num, err := getBalanceFromContract(player, token, instance)
	if err != nil {
		logger.WatchLog.Error().Err(err).Msg("get balance failed")
		return
	}

	balance, _ := utils.NormalizeTokenAmount(num, config.TokenDecimals[config.TokenSymbol[token]]).Float64()

	err = client.PushToGameServer(player, token, balance)
	if err != nil {
		logger.WatchLog.Error().Err(err).Msg("push game server failed")
		return
	}
}

func getBalanceFromContract(player common.Address, token uint8, instance *contract.Baccarat) (*big.Int, error) {
	bal, err := instance.GetBalance0(&bind.CallOpts{From: player}, token)
	if err != nil {
		return big.NewInt(0), err
	}
	return bal, nil
}
