package client

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/si229/web3-api/internal/config"
)

func PushToGameServer(player common.Address, token uint8, balance float64) error {

	symbol := config.TokenSymbol[token]

	reqBody := map[string]interface{}{
		"address": player.Hex(),
		"symbol":  symbol,
		"balance": balance,
	}

	b, _ := json.Marshal(reqBody)

	resp, err := http.Post(
		"http://127.0.0.1:7000/api/balance/update",
		"application/json",
		bytes.NewBuffer(b),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
