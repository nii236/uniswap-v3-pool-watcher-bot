package utils

import (
	"math/big"
)

// Converts hex to decimal
func ConvertHexToDecimal(hexVal string) (*big.Float, error) {
	num := new(big.Float)
	num, _, err := num.Parse(hexVal, 16)
	return num, err
}
