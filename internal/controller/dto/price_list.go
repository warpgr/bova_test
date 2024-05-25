package dto

import "strconv"

type SymbolPrice struct {
	Pair   string `json:"pair"`
	Amount string `json:"amount"`
}

type PriceListDto struct {
	Ltp []SymbolPrice `json:"ltp"`
}

func BuildPriceListDto(priceList map[string]float64) PriceListDto {
	ltp := make([]SymbolPrice, 0, len(priceList))
	for pair, amount := range priceList {
		ltp = append(ltp, SymbolPrice{Pair: pair, Amount: strconv.FormatFloat(amount, 'f', 2, 64)})
	}

	return PriceListDto{
		Ltp: ltp,
	}
}
