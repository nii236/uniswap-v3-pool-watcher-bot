package utils

import (
	"math/big"
)

func CalcUnclaimedFees(id int, val string) string {
	x := ConvertHexToDecimal(val)
	var unclaimedFees *big.Float
	var expo *big.Float
	if id == 0 {
		expo = big.NewFloat(1000000000000000000)
		unclaimedFees = new(big.Float).Quo(x, expo)
	} else {
		expo = big.NewFloat(1000000)
		unclaimedFees = new(big.Float).Quo(x, expo)
	}
	return unclaimedFees.String()
}
