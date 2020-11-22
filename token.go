package bigone

import (
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

func SignAuthenticationToken(key, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"type":       "OpenAPIV2",
		"sub":        key,
		"nonce":      strconv.FormatInt(time.Now().UnixNano(), 10),
		"recvWindow": "150000",
	})
	block := []byte(secret)
	return token.SignedString(block)
}
