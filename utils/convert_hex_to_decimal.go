package utils

import (
	"math/big"
)

func ConvertHexToDecimal(hexVal string) (*big.Float, error) {
	num := new(big.Float)
	num, _, err := num.Parse(hexVal, 16)
	return num, err
}
