package gz

import (
	"bytes"
	"fmt"
	"io"

	"github.com/klauspost/compress/gzip"
)

func Decompress(inputBytes []byte) (string, error) {
	var result bytes.Buffer

	decoded, err := gzip.NewReader(bytes.NewBuffer(inputBytes))
	if err != nil {
		return "", fmt.Errorf("cannot create gzip reader: %w", err)
	}
	defer decoded.Close()

	_, err = io.Copy(&result, decoded)
	if err != nil {
		return "", fmt.Errorf("cannot copy from gzip reader: %w", err)
	}

	return result.String(), nil
}
