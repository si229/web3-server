package blockchain

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

func NewAuth(AccountPrivateKey string, chainID *big.Int) (*bind.TransactOpts, error) {

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(AccountPrivateKey, "0x"))

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	return auth, nil
}
