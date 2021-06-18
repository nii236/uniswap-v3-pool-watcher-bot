package utils

import (
	"log"
	"math/big"
)

func ConvertHexToDecimal(hexVal string) *big.Float {
	num := new(big.Float)
	num, _, err := num.Parse(hexVal, 16)
	if err != nil {
		log.Println(err)
	}
	return num
}
