package web3

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type Artifact struct {
	ABI json.RawMessage `json:"abi"`
}

var ContractAbi abi.ABI

func Init() {
	var err error
	if ContractAbi, err = loadABI("./abi/baccarat.abi"); err != nil {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		log.Fatal("load abi err ", err, wd)
	}
}

func loadABI(path string) (abi.ABI, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return abi.ABI{}, err
	}

	var artifact Artifact

	err = json.Unmarshal(data, &artifact)
	if err != nil {
		return abi.ABI{}, err
	}

	return abi.JSON(
		strings.NewReader(string(artifact.ABI)),
	)
}
