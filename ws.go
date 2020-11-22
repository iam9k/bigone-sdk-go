package bigone

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/golang/protobuf/proto"
	"net"
)

const WsEndpoint = "wss://big.one/ws/v2"

const (
	AuthenticateCustomerRequestId = iota
	SubscribeMarketDepthRequestId
	SubscribeMarketCandlesRequestId
	SubscribeMarketsTickerRequestId
	SubscribeMarketTradesRequestId
	SubscribeViewerAccountsRequestId
	SubscribeViewerOrdersRequestId
)

func WriteClientBinary(conn net.Conn, request *Request) error {
	data, err := proto.Marshal(request)
	if err != nil {
		fmt.Println("proto.Marshal error:", err)
		return err
	}
	return wsutil.WriteClientBinary(conn, data)
}

func CreateWsConn() (conn net.Conn, br *bufio.Reader, hs ws.Handshake, err error) {
	dialer := ws.Dialer{}
	dialer.Protocols = []string{"proto"}
	return dialer.Dial(context.Background(), WsEndpoint)
}

func AuthenticateCustomer(conn net.Conn, key, secret string) error {
	token, err := SignAuthenticationToken(key, secret)
	if err != nil {
		return err
	}
	requestId := fmt.Sprintf("%v", AuthenticateCustomerRequestId)
	request := &Request{RequestId: requestId, Payload: &Request_AuthenticateCustomerRequest{AuthenticateCustomerRequest: &AuthenticateCustomerRequest{Token: fmt.Sprintf("Bearer %v", token)}}}
	return WriteClientBinary(conn, request)
}

// Market Depth
func SubscribeMarketDepth(conn net.Conn, market string) error {
	requestId := fmt.Sprintf("%v-%v", SubscribeMarketDepthRequestId, market)
	request := &Request{RequestId: requestId, Payload: &Request_SubscribeMarketDepthRequest{SubscribeMarketDepthRequest: &SubscribeMarketDepthRequest{Market: market}}}
	return WriteClientBinary(conn, request)
}

// Market Candle
func SubscribeMarketCandle(conn net.Conn, market string, limit int64) error {
	requestId := fmt.Sprintf("%v-%v", SubscribeMarketCandlesRequestId, market)
	request := &Request{RequestId: requestId, Payload: &Request_SubscribeMarketCandlesRequest{SubscribeMarketCandlesRequest: &SubscribeMarketCandlesRequest{Market: market, Period: Candle_MIN5, Limit: limit}}}
	return WriteClientBinary(conn, request)
}

// Market Ticker
func SubscribeMarketsTicker(conn net.Conn, markets []string) error {
	bytes, err := json.Marshal(markets)
	if err != nil {
		return err
	}
	marketsString := string(bytes)
	requestId := fmt.Sprintf("%v-%v", SubscribeMarketsTickerRequestId, marketsString)
	request := &Request{RequestId: requestId, Payload: &Request_SubscribeMarketsTickerRequest{SubscribeMarketsTickerRequest: &SubscribeMarketsTickerRequest{Markets: markets}}}
	return WriteClientBinary(conn, request)
}

// Market Trade
func SubscribeMarketTrade(conn net.Conn, market string, limit int64) error {
	requestId := fmt.Sprintf("%v-%v", SubscribeMarketTradesRequestId, market)
	request := &Request{RequestId: requestId, Payload: &Request_SubscribeMarketTradesRequest{SubscribeMarketTradesRequest: &SubscribeMarketTradesRequest{Market: market, Limit: limit}}}
	return WriteClientBinary(conn, request)
}

// User Account
func SubscribeViewerAccount(conn net.Conn) error {
	requestId := fmt.Sprintf("%v", SubscribeViewerAccountsRequestId)
	request := &Request{RequestId: requestId, Payload: &Request_SubscribeViewerAccountsRequest{SubscribeViewerAccountsRequest: &SubscribeViewerAccountsRequest{}}}
	return WriteClientBinary(conn, request)
}

// User Order
func SubscribeViewerOrder(conn net.Conn, market string) error {
	requestId := fmt.Sprintf("%v-%v", SubscribeViewerOrdersRequestId, market)
	request := &Request{RequestId: requestId, Payload: &Request_SubscribeViewerOrdersRequest{SubscribeViewerOrdersRequest: &SubscribeViewerOrdersRequest{Market: market}}}
	return WriteClientBinary(conn, request)
}
