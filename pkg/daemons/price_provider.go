package daemons

import (
	"context"
	"time"

	"github.com/warpgr/bova_test/pkg/exchanges"
	"github.com/warpgr/bova_test/pkg/store"

	log "github.com/sirupsen/logrus"
)

type PriceProvider interface {
	Run(ctx context.Context, fetchPer time.Duration)
}

func NewPriceProvider(exchange exchanges.Exchange, cache store.KVStorage[string, float64]) PriceProvider {
	return &priceProvider{
		exchange: exchange,
		cache:    cache,
	}
}

type priceProvider struct {
	exchange exchanges.Exchange
	cache    store.KVStorage[string, float64]
}

func (pp *priceProvider) Run(ctx context.Context, fetchPer time.Duration) {
	log.Info("Running price provider")
	t := time.NewTicker(fetchPer)
	provide := func() {
		priceList, err := pp.exchange.GetPriceList(ctx)
		if err != nil {
			log.Errorf("Error occurs when trying to get price list. Maybe rate limiter. Suspending 1m.")
			time.Sleep(time.Minute)
		}
		if err = pp.cache.StoreMany(priceList); err != nil {
			log.Errorf("Can't save price list. Error: %v.", err)
		}
	}
	provide()
	for {
		select {
		case <-ctx.Done():
			log.Warnf("Shutting down price fetcher daemon.")
		case <-t.C:
			provide()
		}
	}
}
