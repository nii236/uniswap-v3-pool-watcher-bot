package utils

import "testing"

func TestUnclaimedFees(t *testing.T) {
	calculated_token_0 := CalcUnclaimedFees(0, "DE0B6B3A7640000")
	calculated_token_1 := CalcUnclaimedFees(1, "F4240")
	if calculated_token_0 != "1" {
		t.Errorf("Error for token 0: Expected %s Calculated %s", "1", calculated_token_0)
	}
	if calculated_token_1 != "1" {
		t.Errorf("Error for token 1: Expected %s Calculated %s", "1", calculated_token_1)
	}
}