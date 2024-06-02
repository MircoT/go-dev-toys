package base64

import (
	"encoding/base64"
	"fmt"
)

func Encode(input string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(input))

	return encoded, nil
}

func Decode(input string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", fmt.Errorf("cannot decode from base64 string: %w", err)
	}

	return string(decoded), nil
}
