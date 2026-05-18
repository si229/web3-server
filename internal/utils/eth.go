package utils

import "math/big"

func ToEth(balanceWei *big.Int) float64 {
	ethValue := new(big.Float).Quo(
		new(big.Float).SetInt(balanceWei),
		big.NewFloat(1e18),
	)

	eth, _ := ethValue.Float64()
	return eth
}

func NormalizeTokenAmount(balance *big.Int, decimals uint8) *big.Float {
	fBalance := new(big.Float).SetInt(balance)

	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)

	fDivisor := new(big.Float).SetInt(divisor)

	return new(big.Float).Quo(fBalance, fDivisor)
}

func ToTokenUnits(amount float64, decimals uint8) *big.Int {
	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)

	fAmount := big.NewFloat(amount)
	fDivisor := new(big.Float).SetInt(divisor)
	fResult := new(big.Float).Mul(fAmount, fDivisor)

	result := new(big.Int)
	fResult.Int(result)
	return result
}
