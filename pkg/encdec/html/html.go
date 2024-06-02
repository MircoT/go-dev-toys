package html

import (
	gohtml "html"
)

func Encode(input string) (string, error) {
	return gohtml.EscapeString(input), nil
}

func Decode(input string) (string, error) {
	return gohtml.UnescapeString(input), nil
}
