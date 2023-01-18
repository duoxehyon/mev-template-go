package uniswap_v2

import (
	"math/big"
)
// "out of" should be 1000 for example and fee is 997 for uniswap v2
func (uniV2 UniswapV2) GetAmountIn(amountOut big.Int, reserveIn big.Int, reserveOut big.Int, OutOf big.Int, Fee big.Int) big.Int {
	if reserveOut.Cmp(big.NewInt(0)) == 0 {
		return *big.NewInt(0)
	}
	numerator := big.NewInt(0).Mul(&reserveIn, &amountOut)
	numerator.Mul(numerator, &OutOf)
	denominator := big.NewInt(0).Sub(&reserveOut, &amountOut)
	denominator.Mul(denominator, &Fee)
	divRes := big.NewInt(0).Div(numerator, denominator)
	return *divRes
}

func (uniV2 UniswapV2) GetAmountOut(amountIn big.Int, reserveIn big.Int, reserveOut big.Int, outOf big.Int, fee big.Int) big.Int {
	if reserveIn.Cmp(big.NewInt(0)) == 0 {
		return *big.NewInt(0)
	}
	amountInWithFee := big.NewInt(0).Mul(&amountIn, &fee)
	numerator := big.NewInt(0).Mul(amountInWithFee, &reserveOut)
	denominator := big.NewInt(0).Mul(&reserveIn, &outOf)
	denominator.Add(denominator, amountInWithFee)
	divRes := big.NewInt(0).Div(numerator, denominator)
	return *divRes
}

