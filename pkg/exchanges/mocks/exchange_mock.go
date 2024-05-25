package mocks

import "context"

type ExchangeMock struct {
	Data map[string]float64
	Err  error
}

func (em *ExchangeMock) GetPriceList(ctx context.Context) (map[string]float64, error) {
	return em.Data, em.Err
}
