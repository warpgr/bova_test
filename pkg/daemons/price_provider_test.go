package daemons_test

import (
	"context"
	"testing"
	"time"

	"github.com/warpgr/bova_test/pkg/daemons"
	"github.com/warpgr/bova_test/pkg/exchanges/common"
	"github.com/warpgr/bova_test/pkg/exchanges/mocks"
	"github.com/warpgr/bova_test/pkg/store"

	"github.com/stretchr/testify/require"
)

func TestPriceProviderRun(t *testing.T) {
	mockedExchange := &mocks.ExchangeMock{
		Data: map[string]float64{
			"BTCUSDT": 65000.21,
			"ETHUSDT": 4500.2,
		},
		Err: nil,
	}

	st := store.NewKVMapStorage[string, float64](2)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	t.Cleanup(func() {
		cancel()
	})
	go daemons.NewPriceProvider(mockedExchange, st).Run(ctx, time.Millisecond)

	time.Sleep(time.Second)
	for i := 0; i < 100; i++ {
		price, err := st.Load("BTCUSDT")
		require.NoError(t, err)
		require.Equal(t, 65000.21, price)

		price, err = st.Load("ETHUSDT")
		require.NoError(t, err)
		require.Equal(t, 4500.2, price)

		// Not found.
		_, err = st.Load("ALGOUSDT")
		require.ErrorIs(t, err, common.ErrLoadElement)
		time.Sleep(time.Millisecond)
	}
}
