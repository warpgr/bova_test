package exchanges

import (
	"context"

	"github.com/warpgr/bova_test/pkg/exchanges/internal"
)

type Exchange interface {
	GetPriceList(ctx context.Context) (map[string]float64, error)
}

func NewKrakenExchange(url string) Exchange {
	return internal.NewKraken(url)
}
