package utils

type PoolParams struct {
	from string
	to   string
	data string
}

// Get these data from individual pools
var RegisteredPools = map[string]PoolParams{
	// DAI/USDC POOL
	"DAI/USDC": {
		"0x11e4857bb9993a50c685a79afad4e6f65d518dda",
		"0xc36442b4a4522e871399cd717abdd847ab11fe88",
		"0xfc6f7865000000000000000000000000000000000000000000000000000000000000000500000000000000000000000011e4857bb9993a50c685a79afad4e6f65d518dda00000000000000000000000000000000ffffffffffffffffffffffffffffffff00000000000000000000000000000000ffffffffffffffffffffffffffffffff",
	},
}
