package utils

var WhitelistTelegramAccountIDs []int = []int{576072597, 533648041}

func IsWhitelistedAccount(accountID int) bool {
	for _, id := range WhitelistTelegramAccountIDs {
		if accountID == id {
			return true
		}
	}
	return false
}
