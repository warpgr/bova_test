package service_test

import (
	"testing"

	"github.com/warpgr/bova_test/internal/repository"
	"github.com/warpgr/bova_test/internal/service"
	"github.com/warpgr/bova_test/pkg/store"

	"github.com/stretchr/testify/require"
)

func TestGetPrices(t *testing.T) {
	st := store.NewKVMapStorage[string, float64](2)
	st.Store("BTCUSDT", 65000.2)
	st.Store("INJUSDT", 1.4)
	st.Store("AVAXUSDT", 100.2)

	svc := service.NewPriceList(repository.NewPriceList(st))

	priceList, err := svc.GetPrices([]string{"BTC/USDT", "INJ/USDT"})
	require.NoError(t, err)
	require.Equal(t, 2, len(priceList))

	priceList, err = svc.GetPrices([]string{"INJ/USDT", "RNDR/USDT"})
	require.NoError(t, err)
	require.Equal(t, 1, len(priceList)) // The list contains only INJ price.

	priceList, err = svc.GetPrices([]string{"ATOM/USDT"})
	require.NoError(t, err)
	require.Equal(t, 0, len(priceList))
}

func TestGetAllPrices(t *testing.T) {
	st := store.NewKVMapStorage[string, float64](2)
	st.Store("BTCUSDT", 65000.2)
	st.Store("INJUSDT", 1.4)
	st.Store("AVAXUSDT", 100.2)

	svc := service.NewPriceList(repository.NewPriceList(st))

	priceList, err := svc.GetAllPrices()
	require.NoError(t, err)
	require.Equal(t, 3, len(priceList))

	require.NoError(t, st.Store("BNBUSDT", 601.4))

	priceList, err = svc.GetAllPrices()
	require.NoError(t, err)
	require.Equal(t, 4, len(priceList))
}
