package jwt

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/golang-jwt/jwt"
)

type JWTOutput struct {
	Header  map[string]interface{}
	Payload jwt.Claims
}

func Decode(input string) (string, error) {
	token, err := jwt.Parse(input, func(token *jwt.Token) (interface{}, error) {
		return []byte{}, fmt.Errorf("no verify")
	})
	if err != nil && err.Error() != "no verify" {
		slog.Error("cannot parse token", "error", err.Error())

		return "", fmt.Errorf("cannot parse token: %w", err)
	}

	jwtOutput := JWTOutput{
		Header:  token.Header,
		Payload: token.Claims,
	}

	res, err := json.MarshalIndent(jwtOutput, "", "  ")
	if err != nil {
		return "", fmt.Errorf("cannot convert header: %w", err)
	}

	return string(res), nil
}
