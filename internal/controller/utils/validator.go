package utils

import "regexp"

// ValidateSymbol validates symbol by regexp.
func ValidateSymbol(symbol string) bool {
	re := regexp.MustCompile(`^[A-Z]{3,5}/[A-Z]{3,5}$`)
	return re.MatchString(symbol)
}
