package exchanges

import "strings"

// ConvertToKrakenSymbol converts symbols BTC/USDT to BTCUSDT.
func ConvertToKrakenSymbol(symbol string) string {
	assetNames := strings.Split(symbol, "/")
	return strings.Join(assetNames, "")
}

// ConvertToKrakenSymbol converts symbols BTC/USDT to BTCUSDT symbols slice.
func ConvertToKrakenSymbols(symbols []string) []string {
	converted := make([]string, 0, len(symbols))
	for _, symbol := range symbols {
		converted = append(converted, ConvertToKrakenSymbol(symbol))
	}
	return converted
}
