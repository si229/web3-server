package cashflow

import (
	"context"
	"log"
	"math/big"

	contract "github.com/si229/web3-api/internal/contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/si229/web3-api/internal/utils"
	"github.com/si229/web3-api/internal/utils/logger"
)

func Deposit(auth *bind.TransactOpts, amount *big.Int, contract *contract.Baccarat) {
	auth.Value = amount
	_, err := contract.Deposit0(auth)
	if err != nil {
		log.Fatal("deposit error:", err)
	} else {
		logger.Log.Info().Msgf("deposit %f", utils.ToEth(amount))
	}

}

func DepositPrizePool(auth *bind.TransactOpts, amount *big.Int, contract *contract.Baccarat) {
	auth.Value = amount
	_, err := contract.DepositPrizePool(auth)
	if err != nil {
		log.Fatal("deposit error:", err)
	} else {
		logger.Log.Info().Msgf("deposit %f", utils.ToEth(amount))
	}

}

func WithdrawPrizePool(auth *bind.TransactOpts, amount *big.Int, contract *contract.Baccarat, client *ethclient.Client) {
	auth.Value = big.NewInt(0)
	tx, err := contract.WithdrawPrizePool(auth, amount)
	if err != nil {
		log.Fatal("withdraw error:", err)
	}
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	logger.Log.Info().Msgf("prize pool receipt.Status:%d Withdraw:%s", receipt.Status, tx.Hash().Hex())
}

func WithdrawAmount(auth *bind.TransactOpts, amount *big.Int, contract *contract.Baccarat, client *ethclient.Client) {
	auth.Value = big.NewInt(0)
	tx, err := contract.Withdraw(auth, amount)
	if err != nil {
		log.Fatal("withdraw error:", err)
	}
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	logger.Log.Info().Msgf("balance receipt.Status:%d Withdraw:%s", receipt.Status, tx.Hash().Hex())
}

func GetBalance(fromAddress common.Address, contract *contract.Baccarat) *big.Int {
	ctx := context.Background()
	balanceWei, err := contract.GetBalance(&bind.CallOpts{Context: ctx, From: fromAddress})
	if err != nil {
		log.Fatal("GetBalance error:", err)
	}
	logger.Log.Info().Msgf("account in contract balance:%f", utils.ToEth(balanceWei))
	return balanceWei
}

func GetPrizePoolBalance(address common.Address, contract *contract.Baccarat) *big.Int {
	ctx := context.Background()
	balanceWei, err := contract.GetPrizePool(&bind.CallOpts{Context: ctx, From: address}, common.Address{})
	if err != nil {
		log.Fatal("GetBalance error:", err)
	}
	logger.Log.Info().Msgf("prize pool balance:%f", utils.ToEth(balanceWei))
	return balanceWei
}

func GetBalanceByAddress(fromAddress common.Address, client *ethclient.Client) {
	balanceWei, err := client.BalanceAt(
		context.Background(),
		fromAddress,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	logger.Log.Info().Msgf("account of balance:%f", utils.ToEth(balanceWei))
}

func GetBalanceByContract(contractAddress common.Address, client *ethclient.Client) *big.Int {
	balanceWei, err := client.BalanceAt(
		context.Background(),
		contractAddress,
		nil, // latest block
	)
	if err != nil {
		log.Fatal(err)
	}
	logger.Log.Info().Msgf("contract balance:%f", utils.ToEth(balanceWei))
	return balanceWei
}
