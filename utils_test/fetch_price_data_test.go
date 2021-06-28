package utils_test

import (
	"testing"
	"uniswap-v3-pool-watcher-bot/utils"
)

func TestFetchTokenPrice(t *testing.T) {
	slug := "ethereum"
	_, err := utils.FetchTokenPrice(slug)
	if err != nil {
		t.Error(err)
	}
}
