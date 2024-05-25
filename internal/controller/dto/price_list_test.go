package dto_test

import (
	"strconv"
	"testing"

	"github.com/warpgr/bova_test/internal/controller/dto"

	"github.com/stretchr/testify/require"
)

func TestBuildPriceListDto(t *testing.T) {
	priceList := map[string]float64{
		"BTCUSDT": 65450.3,
		"ETHUSDT": 40000.4,
	}
	built := dto.BuildPriceListDto(priceList)

	for _, priceData := range built.Ltp {
		price, exists := priceList[priceData.Pair]
		require.True(t, exists)
		require.Equal(t, strconv.FormatFloat(price, 'f', 2, 64), priceData.Amount)
	}
}
