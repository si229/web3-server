package blockchain

import (
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	wsClient   *ethclient.Client
	httpClient *ethclient.Client

	wsOnce   sync.Once
	httpOnce sync.Once
)

// WS（监听用）
func InitWSClient(url string) {
	wsOnce.Do(func() {
		c, err := ethclient.Dial(url)
		if err != nil {
			log.Fatalf("ws client error: %v", err)
		}
		wsClient = c
	})
}

func GetWSClient() *ethclient.Client {
	return wsClient
}

// HTTP（查询用）
func InitHTTPClient(url string) {
	httpOnce.Do(func() {
		c, err := ethclient.Dial(url)
		if err != nil {
			log.Fatalf("http client error: %v", err)
		}
		httpClient = c
	})
}

func GetHTTPClient() *ethclient.Client {
	return httpClient
}
