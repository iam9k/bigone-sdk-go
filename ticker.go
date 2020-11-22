package bigone

import (
	"fmt"
	"context"
)

//type Ticker struct {
//	AssetPairName	string `json:asset_pair_name,omitempty"`
//	Bid	PriceLevel `json:bid,omitempty"`
//	Ask	PriceLevel `json:ask,omitempty"`
//	Open	string `json:open,omitempty"`
//	Close	string `json:close,omitempty"`
//	High	string `json:high,omitempty"`
//	Low	string `json:low,omitempty"`
//	Volume	string `json:volume,omitempty"`
//	DailyChange	string `json:daily_change,omitempty"`
//}

func ReadTicker(assetPairName string) (*Ticker, error)  {

	if assetPairName == "" {
		return nil, fmt.Errorf("assetPairName cannot be an empty string!")
	}

	resp, err := HttpRequest(context.Background()).Get(fmt.Sprintf("/asset_pairs/%v/ticker", assetPairName))

	if err != nil {
		return nil, err
	}

	ticker := &Ticker{}

	if err := UnmarshalResponse(resp, ticker); err != nil {
		return nil, err
	}

	return ticker, nil
}

func ReadTickers(assetPairName string) ([]*Ticker, error)  {

	if assetPairName == "" {
		return nil, fmt.Errorf("assetPairName cannot be an empty string!")
	}

	resp, err := HttpRequest(context.Background()).Get(fmt.Sprintf("/asset_pairs/%v/tickers", assetPairName))

	if err != nil {
		return nil, err
	}

	tickers := []*Ticker{}

	if err := UnmarshalResponse(resp, tickers); err != nil {
		return nil, err
	}

	return tickers, nil
}