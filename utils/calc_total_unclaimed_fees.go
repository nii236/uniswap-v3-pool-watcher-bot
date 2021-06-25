package utils

import (
	"log"
	"math/big"
	"strconv"
)

var tokenSymbolToName = map[string]string{
	"WETH": "ethereum",
	"DAI":  "dai",
	"UNI":  "uniswap",
}

// calculates total unclaimed fees for the pool
func CalcTotalUnclaimedFees(
	token0_unclaimed_fees string,
	token1_unclaimed_fees string,
	token0_name string,
	token1_name string,
) (string, error) {
	// Convert unclaimed fees for each token to *big.Float
	token0_float := new(big.Float)
	token1_float := new(big.Float)
	token0_float.SetString(token0_unclaimed_fees)
	token1_float.SetString(token1_unclaimed_fees)

	// Fetch current prices of each token
	token0PriceInDollars, err := FetchTokenPrice(tokenSymbolToName[token0_name])
	if err != nil {
		log.Println(err)
		return "", err
	}
	token1PriceInDollars, err := FetchTokenPrice(tokenSymbolToName[token1_name])
	if err != nil {
		log.Println(err)
		return "", err
	}

	// Now calculate total unclaimed fees for each token
	token0FeesInDollars := new(big.Float)
	token0UnclaimedFees, err := strconv.ParseFloat(token0_unclaimed_fees, 64)
	if err != nil {
		log.Println("Error converting string to float", err)
		return "", err
	}
	token0FeesInDollars.Mul(big.NewFloat(token0PriceInDollars), big.NewFloat(token0UnclaimedFees))

	token1FeesInDollars := new(big.Float)
	token1UnclaimedFees, err := strconv.ParseFloat(token1_unclaimed_fees, 64)
	if err != nil {
		log.Println("Error converting string to float", err)
		return "", err
	}
	token1FeesInDollars.Mul(big.NewFloat(token1PriceInDollars), big.NewFloat(token1UnclaimedFees))

	// Add above values to get total fees
	totalFees := new(big.Float)
	totalFees.Add(token0FeesInDollars, token1FeesInDollars)
	return totalFees.String(), nil
}
