package exchanges_test

import (
	"testing"

	"github.com/warpgr/bova_test/pkg/exchanges"

	"github.com/stretchr/testify/require"
)

func TestKrakenSymbolsConverters(t *testing.T) {
	symbols := []string{
		"BTC/USDT",
		"ETH/USDT",
	}

	convertedSymbol := exchanges.ConvertToKrakenSymbol(symbols[0])
	require.Equal(t, "BTCUSDT", convertedSymbol)

	convertedSymbols := exchanges.ConvertToKrakenSymbols(symbols)
	require.Equal(t, "BTCUSDT", convertedSymbols[0])
	require.Equal(t, "ETHUSDT", convertedSymbols[1])
}
