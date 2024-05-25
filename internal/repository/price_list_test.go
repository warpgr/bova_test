package repository_test

import (
	"testing"

	"github.com/warpgr/bova_test/internal/repository"
	"github.com/warpgr/bova_test/pkg/store"

	"github.com/stretchr/testify/require"
)

func TestGetPrices(t *testing.T) {
	st := store.NewKVMapStorage[string, float64](2)
	st.Store("BTCUSDT", 65000.2)
	st.Store("ETHUSDT", 4000.5)

	rp := repository.NewPriceList(st)

	priceList, err := rp.GetPrices([]string{"BTCUSDT", "ETHUSDT"})
	require.NoError(t, err)
	require.Equal(t, 2, len(priceList))
}

func TestGetAllPrices(t *testing.T) {
	st := store.NewKVMapStorage[string, float64](2)
	st.Store("BTCUSDT", 65000.2)
	st.Store("ETHUSDT", 4000.5)

	rp := repository.NewPriceList(st)
	priceList, err := rp.GetAllPrices()
	require.NoError(t, err)
	require.Equal(t, 2, len(priceList))
}
