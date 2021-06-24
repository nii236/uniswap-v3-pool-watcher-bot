package utils

import (
	"math/big"
)

// Calculates unclaimed fees for token in a pool
// Params: id -> Possible values are {0,1}. 
// id == 0: Refers to token0. id == 1: Refers to token 1
// val: Refers to unscaled hex value of fees
func CalcUnclaimedFees(id int, val string) (string, error) {
	x, err := ConvertHexToDecimal(val)
	if err != nil {
		return "", err;
	}
	expo := big.NewFloat(1000000000000000000)	// scale decimal
	if id == 0 {
		unclaimedFees := new(big.Float).Quo(x, expo)
		return unclaimedFees.String(), nil
	}
	unclaimedFees := new(big.Float).Quo(x, expo)
	return unclaimedFees.String(), nil
}
