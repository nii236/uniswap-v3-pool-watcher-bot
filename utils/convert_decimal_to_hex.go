package utils

import (
	"fmt"
)

func DecimalToHex(dec int) string {
	return fmt.Sprintf("%x", dec)
}
