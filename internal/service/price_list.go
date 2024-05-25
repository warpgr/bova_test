package service

import (
	"github.com/warpgr/bova_test/internal/repository"
	"github.com/warpgr/bova_test/pkg/exchanges"
)

type PriceList interface {
	GetPrices(symbols []string) (map[string]float64, error)
	GetAllPrices() (map[string]float64, error)
}

func NewPriceList(db repository.PriceList) PriceList {
	return &priceList{
		db: db,
	}
}

type priceList struct {
	db repository.PriceList
}

func (svc *priceList) GetPrices(symbols []string) (map[string]float64, error) {
	return svc.db.GetPrices(exchanges.ConvertToKrakenSymbols(symbols))
}

func (svc *priceList) GetAllPrices() (map[string]float64, error) {
	return svc.db.GetAllPrices()
}
