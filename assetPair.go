package bigone

import (
	"context"
)

type Asset struct {
	Id     string `json:id,omitempty"`
	Symbol string `json:"symbol,omitempty"`
	Name   string `json:"name,omitempty"`
}

type AssetPair struct {
	Id             string `json:id,omitempty"`
	QuoteScale     int    `json:quote_scale,omitempty"`
	QuoteAsset     Asset  `json:quote_asset,omitempty"`
	Name           string `json:name,omitempty"`
	BaseScale      int    `json:base_scale,omitempty"`
	MinQuote_value string `json:min_quote_value,omitempty"`
	BaseAsset      Asset  `json:base_asset,omitempty"`
}

func ReadAssetPair() ([]*AssetPair, error) {
	resp, err := HttpRequest(context.Background()).Get("/asset_pairs")

	if err != nil {
		return nil, err
	}

	assetPairs := []*AssetPair{}

	if err := UnmarshalResponse(resp, assetPairs); err != nil {
		return nil, err
	}

	return assetPairs, nil
}
