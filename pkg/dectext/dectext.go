package dectext

import (
	"fmt"
	"os"

	"github.com/h2non/filetype"
)

const (
	numHeaderBytes = 261
)

func GetCompressType(targetFile string) (string, error) {
	// Open a file descriptor
	curFile, _ := os.Open(targetFile)

	// We only have to pass the file header = first 261 bytes
	head := make([]byte, numHeaderBytes)
	if _, err := curFile.Read(head); err != nil {
		return "", fmt.Errorf("cannot read header")
	}

	fileType, err := filetype.Get(head)
	if err != nil {
		return "", fmt.Errorf("cannot get type from header")
	}

	return fileType.Extension, nil
}
