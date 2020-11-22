package bigone

import (
"fmt"
"context"
)

//type Candle struct {
//	Close	string `json:close,omitempty"`
//	High	PriceLevel `json:high,omitempty"`
//	Low	PriceLevel `json:low,omitempty"`
//	Open	string `json:open,omitempty"`
//	Volume	string `json:volume,omitempty"`
//	Time	string `json:time,omitempty"`
//}

func ReadCandles(assetPairName string) ([]*Candle, error)  {

	if assetPairName == "" {
		return nil, fmt.Errorf("assetPairName cannot be an empty string!")
	}

	resp, err := HttpRequest(context.Background()).Get(fmt.Sprintf("/asset_pairs/%v/ticker", assetPairName))

	if err != nil {
		return nil, err
	}

	candles := []*Candle{}

	if err := UnmarshalResponse(resp, candles); err != nil {
		return nil, err
	}

	return candles, nil
}