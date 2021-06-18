package utils

import (
	"math/big"
	"testing"
)

func TestHexToDecimal(t *testing.T) {
	str := "41"
	expected := big.NewInt(65).String()
	calculated := ConvertHexToDecimal(str).String()
	if expected != calculated {
		t.Error("Hex to decimal converter not working properly")
	}
}
