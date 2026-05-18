package router

import (
	"github.com/gin-gonic/gin"
	"github.com/si229/web3-api/internal/config"
	"github.com/si229/web3-api/internal/handlers"
)

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {

	handlerMap := map[string]gin.HandlerFunc{
		"verifySignature":    handlers.VerifySignature,
		"genVerifyParam":     handlers.GenVerifyParam,
		"deposit":            handlers.Deposit,
		"getBalance":         handlers.GetBalance,
		"getContractBalance": handlers.GetContractBalance,
		"addressValidate":    handlers.AddressValidate,
		"withdraw":           handlers.Withdraw,
		"getDepositData":     handlers.GetDepositData,
	}

	for _, route := range cfg.Routes {
		h, ok := handlerMap[route.Handler]
		if !ok {
			panic("unknown handler: " + route.Handler)
		}

		switch route.Method {
		case "GET":
			r.GET(route.Path, h)
		case "POST":
			r.POST(route.Path, h)
		case "PUT":
			r.PUT(route.Path, h)
		case "DELETE":
			r.DELETE(route.Path, h)
		}
	}
}
