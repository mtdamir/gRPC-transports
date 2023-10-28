package main

import (
	"context"
	"fmt"
	"time"
)

// PriceFetcher is an  interface that can fetch a price
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

// PriceFetcher implements  the PriceFetcher interface
type priceFetcher struct {
}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)
}

var PriceMocks = map[string]float64{
	"BTC": 20_000.0,
	"ETH": 200.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	// mimic the http roundtrip
	time.Sleep(100 * time.Millisecond)
	price, ok := PriceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}

	return price, nil
}
