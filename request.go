package bigone

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

const (
	Endpoint = "https://big.one/api/v3"
)

type HttpError struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"message,omitempty"`
}

func (err *HttpError) Error() string {
	return fmt.Sprintf("[%d] %s", err.Code, err.Msg)
}

var httpClient = resty.New().
	SetHeader("Content-Type", "application/json").
	SetHostURL(Endpoint).
	SetTimeout(2 * time.Second)

func HttpRequest(ctx context.Context) *resty.Request {
	return httpClient.R().SetContext(ctx)
}

func DecodeResponse(resp *resty.Response) ([]byte, error) {
	var body struct {
		HttpError
		Data      json.RawMessage `json:"data,omitempty"`
		Code      int             `json:"code,omitempty"`
		Message   string          `json:"message,omitempty"`
		PageToken string          `json:"page_token,omitempty"`
	}

	if err := json.Unmarshal(resp.Body(), &body); err != nil {
		if resp.IsError() {
			return nil, &HttpError{
				Code: resp.StatusCode(),
				Msg:  resp.Status(),
			}
		}

		return nil, err
	}

	if body.Data == nil {
		return nil, &HttpError{
			Code: -1,
			Msg:  body.Message,
		}
	}

	if body.HttpError.Code > 0 {
		return nil, &body.HttpError
	}

	return body.Data, nil
}

func UnmarshalResponse(resp *resty.Response, v interface{}) error {
	data, err := DecodeResponse(resp)
	if err != nil {
		return err
	}

	if v != nil {
		return json.Unmarshal(data, v)
	}

	return nil
}
