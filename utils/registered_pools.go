package utils

import "math/big"

type PoolParams struct {
	from            string
	to              string
	data            string
	lastTrackedFees *big.Float // stores last tracked total unclaimed fees in the pool
}

// Get these data from individual pools - More in READMW
// Represents the pools which we wish to track
var RegisteredPools = map[string]*PoolParams{
	"UNI/WETH": { // https://app.uniswap.org/#/pool/1
		"0x11e4857bb9993a50c685a79afad4e6f65d518dda",
		"0xc36442b4a4522e871399cd717abdd847ab11fe88",
		"0xfc6f7865000000000000000000000000000000000000000000000000000000000000000100000000000000000000000011e4857bb9993a50c685a79afad4e6f65d518dda00000000000000000000000000000000ffffffffffffffffffffffffffffffff00000000000000000000000000000000ffffffffffffffffffffffffffffffff",
		big.NewFloat(0),
	},
	"DAI/WETH": { // https://app.uniswap.org/#/pool/25
		"0xdd0d6c26a03d6f6541471d44179f56d478f50f6b",
		"0xc36442b4a4522e871399cd717abdd847ab11fe88",
		"0xfc6f78650000000000000000000000000000000000000000000000000000000000000019000000000000000000000000dd0d6c26a03d6f6541471d44179f56d478f50f6b00000000000000000000000000000000ffffffffffffffffffffffffffffffff00000000000000000000000000000000ffffffffffffffffffffffffffffffff",
		big.NewFloat(0),
	},
}
