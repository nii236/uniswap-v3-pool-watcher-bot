package utils

import (
	"fmt"
	"log"
	"math/big"
	"strings"
)

// This function chooses only those pools:
// 1. which have crossed the threshold for unclaimed fees
// 2. whose last tracked value was less than the threshold
func PoolCrossedThreshold(threshold float64, totalFeesInFloat *big.Float, token0_name string, token1_name string) bool {
	pairName := fmt.Sprintf("%s/%s", token0_name, token1_name)
	lastValue := RegisteredPools[pairName].lastTrackedFees
	currentValue := totalFeesInFloat

	// Update the lastTrackedFees of the pool with the currentValue
	RegisteredPools[pairName].lastTrackedFees = currentValue
	return big.NewFloat(threshold).Cmp(currentValue) <= 0 && // currentVal > threshold
		big.NewFloat(threshold).Cmp(lastValue) >= 0 // lastVal < threshold
}

func FilterThresholdCrossingPools(gethUrl string, threshold float64) []string {
	thresholdPools := make([]string, 0)
	for k, v := range RegisteredPools {
		token0_name := strings.Split(k, "/")[0]
		token1_name := strings.Split(k, "/")[1]
		uniswapResp, err := UniswapAPICall(v.from, v.to, v.data, gethUrl)
		if err != nil {
			log.Printf("Error %v", err)
			continue
		}

		//get token0 unclaimed fees
		token0_unclaimed_fees, err := CalcUnclaimedFees(0, uniswapResp[2:66])
		if err != nil {
			log.Printf("Error %v", err)
			continue
		}

		//get token1 unclaimed fees
		token1_unclaimed_fees, err := CalcUnclaimedFees(1, uniswapResp[66:])
		if err != nil {
			log.Printf("Error %v", err)
			continue
		}

		// Now get total unclaimed fees in dollars
		total_unclaimed_fees, err := CalcTotalUnclaimedFees(token0_unclaimed_fees, token1_unclaimed_fees, token0_name, token1_name)
		totalFeesInFloat := new(big.Float)
		totalFeesInFloat.SetString(total_unclaimed_fees)

		if PoolCrossedThreshold(threshold, totalFeesInFloat, token0_name, token1_name) == true {
			thresholdPools = append(thresholdPools, fmt.Sprintf("%s/%s: %s", token0_name, token1_name, total_unclaimed_fees))
		}

		if err != nil {
			log.Println(err)
			continue
		}
	}
	return thresholdPools
}
