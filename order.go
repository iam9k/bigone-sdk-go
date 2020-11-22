package bigone

import (
	"context"
	"fmt"
)

//type Order struct {
//	Id                int64  `json:id,omitempty"`
//	AssetPairName     string `json:asset_pair_name,omitempty"`
//	Price             string `json:price,omitempty"`
//	Amount            string `json:amount,omitempty"`
//	FilledAmount      string `json:filled_amount,omitempty"`
//	AvgDealPrice      string `json:avg_deal_price,omitempty"`
//	Side              string `json:side,omitempty"`
//	State             string `json:state,omitempty"`
//	CreatedAt         string `json:created_at,omitempty"`
//	UpdatedAt         string `json:updated_at,omitempty"`
//	Type              string `json:type,omitempty"`
//	StopPrice         string `json:stop_price,omitempty"`
//	Operator          string `json:operator,omitempty"`
//	ImmediateOrCancel bool   `json:immediate_or_cancel,omitempty"`
//	PostOnly          bool   `json:post_only,omitempty"`
//}

func ReadOrders(token string) ([]*Order, error) {
	resp, err := HttpRequest(context.Background()).SetAuthToken(token).Get("/viewer/orders")

	if err != nil {
		return nil, err
	}

	orders := []*Order{}

	if err := UnmarshalResponse(resp, orders); err != nil {
		return nil, err
	}

	return orders, nil
}

func ReadOrder(orderId int64, token string) (*Order, error) {
	if token == "" {
		return nil, fmt.Errorf("token cannot be an empty string!")
	}

	resp, err := HttpRequest(context.Background()).SetAuthToken(token).Get(fmt.Sprintf("/viewer/orders/%v", orderId))

	if err != nil {
		return nil, err
	}

	order := &Order{}

	if err := UnmarshalResponse(resp, order); err != nil {
		return nil, err
	}

	return order, nil
}

//TODO
func createOrder() {

}

func createMultiOrder() {

}

func cancelOrder() {

}

func cancelAllOrders() {

}
