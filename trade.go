package bigone

import (
	"context"
	"fmt"
)

//type Trade struct {
//	Id            int64  `json:id,omitempty"`
//	AssetPairName string `json:asset_pair_name,omitempty"`
//	Price         string `json:price,omitempty"`
//	Amount        string `json:amount,omitempty"`
//	TakerSide     string `json:taker_side,omitempty"`
//	CreatedAt     string `json:created_at,omitempty"`
//}

type UserTrade struct {
	Trade
}

func ReadTrades(assetPairName string) (*Trade, error) {

	if assetPairName == "" {
		return nil, fmt.Errorf("assetPairName cannot be an empty string!")
	}

	resp, err := HttpRequest(context.Background()).Get(fmt.Sprintf("/asset_pairs/%v/trades", assetPairName))

	if err != nil {
		return nil, err
	}

	trade := &Trade{}

	if err := UnmarshalResponse(resp, trade); err != nil {
		return nil, err
	}

	return trade, nil
}

type ReadUserTradesParams struct {
	AssetPairName string
	PageToken string
	limit string
}

func ReadUserTrades( params ReadUserTradesParams, token string) ([]*UserTrade, error) {
	values := make(map[string][]string)
	values["AssetPairName"] = []string{params.AssetPairName}
	values["PageToken"] = []string{params.PageToken}
	values["limit"] = []string{params.limit}

	resp, err := HttpRequest(context.Background()).SetQueryParamsFromValues(values).Get("/viewer/trades")

	if err != nil {
		return nil, err
	}

	userTrades := []*UserTrade{}

	if err := UnmarshalResponse(resp, userTrades); err != nil {
		return nil, err
	}

	return userTrades, nil
}
