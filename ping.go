package bigone

import (
	"context"
	"fmt"
)

func PingServer() (int64, error)  {
	resp, err := HttpRequest(context.Background()).Get("/ping")

	if err != nil {
		return 0, err
	}

	var data struct {
		Timestamp int64 `json:"timestamp,omitempty"`
	}

	fmt.Println("resp", resp)

	if err := UnmarshalResponse(resp, &data); err != nil {
		return 0, err
	}

	return data.Timestamp, nil
}