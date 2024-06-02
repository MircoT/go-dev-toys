package url

import (
	"fmt"
	gourl "net/url"
)

func Encode(input string) (string, error) {
	return gourl.QueryEscape(input), nil
}

func Decode(input string) (string, error) {
	decoded, err := gourl.QueryUnescape(input)
	if err != nil {
		return "", fmt.Errorf("cannot decode from url string: %w", err)
	}

	return decoded, nil
}
