package blockchain

import (
	"log"
	"sync"

	contract "github.com/si229/web3-api/internal/contract"

	"github.com/ethereum/go-ethereum/common"
)

// 合约结构（WS / HTTP 分离）
type BaccaratContracts struct {
	WS   *contract.Baccarat
	HTTP *contract.Baccarat
}

var (
	baccarat     BaccaratContracts
	baccaratOnce sync.Once
)

// 初始化合约
func InitBaccarat(addr string) {
	baccaratOnce.Do(func() {
		ws := GetWSClient()
		http := GetHTTPClient()

		wsInstance, err := contract.NewBaccarat(
			common.HexToAddress(addr),
			ws,
		)
		if err != nil {
			log.Fatalf("init WS contract failed: %v", err)
		}

		httpInstance, err := contract.NewBaccarat(
			common.HexToAddress(addr),
			http,
		)
		if err != nil {
			log.Fatalf("init HTTP contract failed: %v", err)
		}

		baccarat = BaccaratContracts{
			WS:   wsInstance,
			HTTP: httpInstance,
		}
	})
}

// 获取合约
func GetBaccarat() BaccaratContracts {
	if baccarat.WS == nil || baccarat.HTTP == nil {
		log.Fatal("baccarat contract not initialized")
	}
	return baccarat
}
