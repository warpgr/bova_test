package utils_test

import (
	"testing"

	"github.com/warpgr/bova_test/internal/controller/utils"

	"github.com/stretchr/testify/require"
)

func TestValidateSymbol(t *testing.T) {
	require.False(t, utils.ValidateSymbol("BTCUSDT"))
	require.True(t, utils.ValidateSymbol("BTC/USDT"))
	require.False(t, utils.ValidateSymbol("SomethingWrong/"))
	require.False(t, utils.ValidateSymbol("$QSAklds"))
}
