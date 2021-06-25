package utils

import (
	"fmt"
	"log"
	"math/big"
	"strings"
)

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
		if big.NewFloat(threshold).Cmp(totalFeesInFloat) <= 0 {
			thresholdPools = append(thresholdPools, fmt.Sprintf("%s/%s: %s", token0_name, token1_name, total_unclaimed_fees))
		}
		
		if err != nil {
			log.Println(err)
			continue
		}
	}
	return thresholdPools
}