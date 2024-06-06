package zip

import (
	"bytes"
	"fmt"
	"io"

	"github.com/klauspost/compress/zip"
)

type compressedReaderAt struct {
	data []byte
	R    io.Reader
	N    int64
}

func (u *compressedReaderAt) ReadAt(p []byte, off int64) (n int, err error) {
	if off < u.N || off >= int64(len(u.data)) {
		return 0, fmt.Errorf("offset %d is not valid", off)
	}

	diff := off - u.N
	written, err := io.CopyN(io.Discard, u.R, diff)
	u.N += written

	if err != nil {
		return int(written), err
	}

	n, err = u.R.Read(p)
	// u.N += int64(n)

	u.R = bytes.NewBuffer(u.data)
	u.N = 0

	return
}

func newCompressedReaderAt(data []byte) io.ReaderAt {
	return &compressedReaderAt{data: data, R: bytes.NewBuffer(data)}
}

func Decompress(inputBytes []byte) (string, error) {
	var result bytes.Buffer

	zipContainer, err := zip.NewReader(
		newCompressedReaderAt(inputBytes),
		int64(len(inputBytes)),
	)
	if err != nil {
		return "", fmt.Errorf("cannot create zip container reader: %w", err)
	}

	decoded, err := zipContainer.File[0].Open()
	if err != nil {
		return "", fmt.Errorf("cannot create zip reader: %w", err)
	}
	defer decoded.Close()

	_, err = io.Copy(&result, decoded)
	if err != nil {
		return "", fmt.Errorf("cannot copy from zip reader: %w", err)
	}

	return result.String(), nil
}
