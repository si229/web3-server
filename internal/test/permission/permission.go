package permission

import (
	"log"

	contract "github.com/si229/web3-api/internal/contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/si229/web3-api/internal/utils/logger"
)

func IsOwner(address string, contract *contract.Baccarat) bool {

	isOwner, err := contract.IsOwner(&bind.CallOpts{From: common.HexToAddress(address)})

	if err != nil {
		log.Fatal(" error:", err)
	} else {
		logger.Log.Info().Msgf(" %t", isOwner)
	}
	return isOwner
}

func TransferOwnership(auth *bind.TransactOpts, address string, contract *contract.Baccarat) {
	_, err := contract.TransferOwnership(auth, common.HexToAddress(address))

	if err != nil {
		log.Fatal(" error:", err)
	} else {
		logger.Log.Info().Msgf("new owner: %s", address)
	}
}
