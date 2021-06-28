package utils

var WhitelistTelegramAccountIDs []int

//Returns whether the given account is one of the whitelisted telegram account ids
func IsWhitelistedAccount(accountID int) bool {
	// Assign the account ids to the global variable WhitelistedTelegramAccountIDs
	// so that it can be used across all functions in the given package
	for _, id := range WhitelistTelegramAccountIDs {
		if accountID == id {
			return true
		}
	}
	return false
}
