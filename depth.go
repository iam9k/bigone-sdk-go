package bigone

import (
	"fmt"
	"context"
)


//type PriceLevel struct {
//	Price string  `json:price,omitempty"`
//	Quantity string `json:"quantity,omitempty"`
//	OrderCount int64 `json:"order_count,omitempty"`
//}


//type Depth struct {
//	AssetPairName string `json:"asset_pair_name,omitempty"`
//	Bids []PriceLevel `json:"bids,omitempty"`
//	Asks []PriceLevel `json:"asks,omitempty"`
//}


/**
Order Book is the ask orders and bid orders collection of a asset pair
*/
func ReadDepth(assetPairName, limit  string) (*Depth, error)  {

	if assetPairName == "" || limit == "" {
		return nil, fmt.Errorf("assetPairName、limit cannot be an empty string!")
	}

	resp, err := HttpRequest(context.Background()).SetQueryParam("limit", limit).Get(fmt.Sprintf("/asset_pairs/%v/depth", assetPairName))

	if err != nil {
		return nil, err
	}

	depths := &Depth{}

	if err := UnmarshalResponse(resp, depths); err != nil {
		return nil, err
	}

	return depths, nil
}