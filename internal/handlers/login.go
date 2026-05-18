package handlers

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	request "github.com/si229/web3-api/internal/model"
	signature "github.com/si229/web3-api/internal/security"
	"github.com/si229/web3-api/internal/utils/logger"
)

func AddressValidate(c *gin.Context) {

	var req request.VerifyValidate

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	b := IsValidAddress(req.Address)

	c.JSON(200, gin.H{"data": b})

}

func VerifySignature(c *gin.Context) {

	var req request.VerifySignatureReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	logger.Log.Info().Msgf("body:%+v", req)

	b := signature.Verify(req.Address, req.Once, req.Signature)

	logger.Log.Info().Msgf("body:%t", b)

	c.JSON(200, gin.H{"data": b})
}

func GenVerifyParam(c *gin.Context) {
	var req request.GenVerifyParamReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	address, sign := signature.PersonalSign("df57089febbacf7ba0bc227dafbffa9fc08a93fdc68e1e42411a14efcf23656e", req.Once)
	s := signature.Verify(address, req.Once, sign)
	c.JSON(200, gin.H{"signature": sign, "address": address, "once": req.Once, "b": s})
}

func IsValidAddress(addr string) bool {
	// 必须是 0x 开头
	if !strings.HasPrefix(addr, "0x") {
		return false
	}

	// 长度必须正确
	if len(addr) != 42 {
		return false
	}

	// common.IsHexAddress 会做基础 hex 校验
	if !common.IsHexAddress(addr) {
		return false
	}

	// 解析地址（会标准化）
	address := common.HexToAddress(addr)

	// 如果输入是 checksum 格式，做严格校验
	// 如果全小写/全大写，认为是未 checksum 地址（允许）
	if hasMixedCase(addr) {
		return common.HexToAddress(addr).Hex() == address.Hex()
	}

	return true
}

// 判断是否混合大小写（EIP-55 checksum 特征）
func hasMixedCase(s string) bool {
	hasLower := false
	hasUpper := false

	for _, c := range s {
		if c >= 'a' && c <= 'f' {
			hasLower = true
		}
		if c >= 'A' && c <= 'F' {
			hasUpper = true
		}
	}
	return hasLower && hasUpper
}
