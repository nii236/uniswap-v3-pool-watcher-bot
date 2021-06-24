package utils_test

import (
	"math/big"
	"testing"
	"uniswap-v3-pool-watcher-bot/utils"
)

func TestHexToDecimal(t *testing.T) {
	str := "41"
	expected := big.NewInt(65).String()
	calculated, err := utils.ConvertHexToDecimal(str)
	if err != nil {
		t.Error("Error: ", err)
	}
	if expected != calculated.String() {
		t.Error("Hex to decimal converter not working properly")
	}
}
