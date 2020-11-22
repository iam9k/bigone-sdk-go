package bigone

import (
	"context"
	"fmt"
)


//type Account struct {
//	AssetSymbol	string `json:asset_symbol,omitempty"`
//	Balance	PriceLevel `json:balance,omitempty"`
//	LockedBalance	PriceLevel `json:locked_balance,omitempty"`
//}

func ReadAccounts(token string) ([]*Account, error)  {
	resp, err := HttpRequest(context.Background()).SetAuthToken(token).Get("/viewer/accounts")

	if err != nil {
		return nil, err
	}

	accounts := []*Account{}

	if err := UnmarshalResponse(resp, accounts); err != nil {
		return nil, err
	}

	return accounts, nil
}

func ReadAccount(asset_symbol , token string) (*Account, error)  {
	if asset_symbol == "" {
		return nil, fmt.Errorf("asset_symbol cannot be an empty string!")
	}

	resp, err := HttpRequest(context.Background()).SetAuthToken(token).Get(fmt.Sprintf("/viewer/accounts/%v", asset_symbol))

	if err != nil {
		return nil, err
	}

	account := &Account{}

	if err := UnmarshalResponse(resp, account); err != nil {
		return nil, err
	}

	return account, nil
}