package utils_test

import (
	"testing"
	"uniswap-v3-pool-watcher-bot/utils"
)

func TestFetchTokenPrice(t *testing.T) {
	slug := "ethereum"
	expected := 2.887
	calculated, err := utils.FetchTokenPrice(slug)
	if err != nil {
		t.Error(err)
	}
	if expected != calculated {
		t.Errorf("Expected %f but got %f", expected, calculated)
	}
}
