package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

type Kraken struct {
	client *http.Client
	url    string
}

func NewKraken(url string) *Kraken {
	return &Kraken{
		client: &http.Client{},
		url:    url,
	}
}

func (k *Kraken) GetPriceList(ctx context.Context) (map[string]float64, error) {
	response, err := k.client.Get(k.url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var result krakenResponse
	if err = json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}
	extracted := make(map[string]float64, len(result.Result))
	for symbol, data := range result.Result {
		price, err := strconv.ParseFloat(data.C[0], 64)
		if err != nil {
			continue
		}
		extracted[symbol] = price
	}
	return extracted, nil
}

type krakenResponse struct {
	Error  interface{} `json:"error"`
	Result map[string]priceParameters
}

type priceParameters struct {
	A []string `json:"a"`
	B []string `json:"b"`
	C []string `json:"c"`
	V []string `json:"v"`
	P []string `json:"p"`
	T []int    `json:"t"`
	L []string `json:"l"`
	H []string `json:"h"`
	O string   `json:"o"`
}
