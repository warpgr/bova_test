package internal_test

import (
	"context"
	"testing"

	"github.com/warpgr/bova_test/pkg/exchanges/internal"

	"github.com/stretchr/testify/require"
)

func TestKraken(t *testing.T) {
	url := "https://api.kraken.com/0/public/Ticker"

	kraken := internal.NewKraken(url)
	priceList, err := kraken.GetPriceList(context.Background())
	require.NoError(t, err)
	require.NotEqual(t, 0, len(priceList))
}
