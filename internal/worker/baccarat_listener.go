package worker

import (
	"context"
	"time"

	"github.com/si229/web3-api/blockchain"
	contract "github.com/si229/web3-api/internal/contract"
	"github.com/si229/web3-api/internal/service"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/si229/web3-api/internal/utils/logger"
)

func StartBaccaratListener() {
	for {
		err := runListener()
		if err != nil {
			logger.WatchLog.Error().Err(err).Msgf("listener stopped, retrying...")
		}

		time.Sleep(3 * time.Second)
	}
}

func runListener() error {
	instance := blockchain.GetBaccarat().WS
	depositLogs := make(chan *contract.BaccaratDeposit)
	withdrawLogs := make(chan *contract.BaccaratWithdraw)
	settleLogs := make(chan *contract.BaccaratSettle)
	depositPrizeLogs := make(chan *contract.BaccaratDepositPrize)
	withdrawPrizeLogs := make(chan *contract.BaccaratWithdrawPrize)

	ctx := context.Background()

	// 分别订阅事件
	subDeposit, err := instance.WatchDeposit(&bind.WatchOpts{Context: ctx}, depositLogs, nil, nil)
	if err != nil {
		return err
	}
	subWithdraw, err := instance.WatchWithdraw(&bind.WatchOpts{Context: ctx}, withdrawLogs, nil, nil)
	if err != nil {
		return err
	}
	subSettle, err := instance.WatchSettle(&bind.WatchOpts{Context: ctx}, settleLogs, nil, nil)
	if err != nil {
		return err
	}
	subDepositPrize, err := instance.WatchDepositPrize(&bind.WatchOpts{Context: ctx}, depositPrizeLogs, nil, nil)
	if err != nil {
		return err
	}
	subWithdrawPrize, err := instance.WatchWithdrawPrize(&bind.WatchOpts{Context: ctx}, withdrawPrizeLogs, nil, nil)
	if err != nil {
		return err
	}
	defer subDeposit.Unsubscribe()
	defer subWithdraw.Unsubscribe()
	defer subSettle.Unsubscribe()
	defer subDepositPrize.Unsubscribe()
	defer subWithdrawPrize.Unsubscribe()

	// 统一处理
	for {
		select {
		case err := <-subDeposit.Err():
			return err
		case err := <-subWithdraw.Err():
			return err
		case err := <-subSettle.Err():
			return err
		case err := <-subDepositPrize.Err():
			return err
		case err := <-subWithdrawPrize.Err():
			return err
		case d := <-depositLogs:
			go service.HandleUpdateBalance(d.Player, d.Token, blockchain.GetBaccarat().HTTP)
			logger.WatchLog.Info().Msgf("Deposit:%s %s %d %d", d.Player, d.Token, d.Amount, d.Balance)
		case w := <-withdrawLogs:
			go service.HandleUpdateBalance(w.Player, w.Token, blockchain.GetBaccarat().HTTP)
			logger.WatchLog.Info().Msgf("Withdraw:%s %s %d %d", w.Player, w.Token, w.Amount, w.Balance)
		case s := <-settleLogs:
			go service.HandleUpdateBalance(s.Player, s.Token, blockchain.GetBaccarat().HTTP)
			logger.WatchLog.Info().Msgf("Settle:%s %s %d %d", s.Player, s.Token, s.Amount, s.Balance)
		case dp := <-depositPrizeLogs:
			logger.WatchLog.Info().Msgf("depositPrize:%s %s %d %d", dp.Player, dp.Token, dp.Amount, dp.Balance)
		case wp := <-withdrawPrizeLogs:
			logger.WatchLog.Info().Msgf("WithdrawPrize:%s %s %d %d", wp.Player, wp.Token, wp.Amount, wp.Balance)
		}
	}

}
