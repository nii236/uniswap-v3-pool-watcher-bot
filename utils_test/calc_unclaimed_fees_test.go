package utils_test

import (
	"testing"
	"uniswap-v3-pool-watcher-bot/utils"
)

func TestUnclaimedFees(t *testing.T) {
	calculated_token_0, err := utils.CalcUnclaimedFees(0, utils.DecimalToHex(1000000000000000000))
	if err != nil {
		t.Errorf("Error for token 0: %v", err)
	}
	calculated_token_1, err := utils.CalcUnclaimedFees(1, utils.DecimalToHex(1000000))
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
