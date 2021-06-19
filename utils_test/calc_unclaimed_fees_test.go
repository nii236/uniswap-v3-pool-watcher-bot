package utils_test

import (
	"fmt"
	"testing"
	"uniswap-v3-pool-watcher-bot/utils"
)

func DecimalToHex (dec int) string {
	return fmt.Sprintf("%x", dec)
}

func TestUnclaimedFees(t *testing.T) {
	calculated_token_0, err := utils.CalcUnclaimedFees(0, DecimalToHex(1000000000000000000))
	if err != nil {
		t.Errorf("Error for token 0: %v", err)
	}
	calculated_token_1, err := utils.CalcUnclaimedFees(1, DecimalToHex(1000000))
	if err != nil {
		t.Errorf("Error for token 1: %v", err)
	}
	if calculated_token_0 != "1" {
		t.Errorf("Error for token 0: Expected %s Calculated %s", "1", calculated_token_0)
	}
	if calculated_token_1 != "1" {
		t.Errorf("Error for token 1: Expected %s Calculated %s", "1", calculated_token_1)
	}
}
