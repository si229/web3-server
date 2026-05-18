package handlers

import (
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/si229/web3-api/internal/config"
	request "github.com/si229/web3-api/internal/model"
	"github.com/si229/web3-api/internal/utils"
	"github.com/si229/web3-api/internal/web3"
)

func GetDepositData(c *gin.Context) {
	var req request.DepositDataReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if token, ok := config.SymbolToken[req.Symbol]; ok {
		depositAmount := utils.ToTokenUnits(req.Amount, config.TokenDecimals[req.Symbol])
		if data, err := web3.ContractAbi.Pack("deposit", token, depositAmount); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"data": "0x" + BytesToHex(data),
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": "Invalid symbol",
		})
	}

}

func BytesToHex(data []byte) string {
	return hex.EncodeToString(data)
}
