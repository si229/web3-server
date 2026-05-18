package repository

import (
	"encoding/json"
	"log"
	"os"
)

type Account struct {
	Index      int    `json:"index"`
	Address    string `json:"address"`
	PrivateKey string `json:"private_key"`
	Balance    string `json:"balance"`
}

var Accounts []Account

func init() {
	data, err := os.ReadFile("./internal/contract/accounts.json")
	if err != nil {
		log.Fatal(err)
	}
	_ = json.Unmarshal(data, &Accounts)
}
