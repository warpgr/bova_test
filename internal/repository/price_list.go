package repository

import "github.com/warpgr/bova_test/pkg/store"

type PriceList interface {
	GetPrices(symbols []string) (map[string]float64, error)
	GetAllPrices() (map[string]float64, error)
}

func NewPriceList(kvStore store.KVStorage[string, float64]) PriceList {
	return &priceList{
		kvStore: kvStore,
	}
}

type priceList struct {
	kvStore store.KVStorage[string, float64]
}

func (rp *priceList) GetPrice(symbol string) (float64, error) {
	return rp.kvStore.Load(symbol)
}

func (rp *priceList) GetPrices(symbols []string) (map[string]float64, error) {
	return rp.kvStore.LoadMany(symbols)
}

func (rp *priceList) GetAllPrices() (map[string]float64, error) {
	return rp.kvStore.LoadAll(), nil
}
