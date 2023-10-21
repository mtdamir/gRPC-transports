package main

import (
	"context"
	"fmt"
)

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct {
}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {

}

var PriceMocks = map[string]float64{
	"BTC": 20_000.0,
	"ETH": 200.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	price, ok := PriceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}

	return price, nil
}
