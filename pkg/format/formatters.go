package format

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

func CheckArgs(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return errors.New("requires the format and the input string")
	} else if len(args) > 2 {
		return errors.New("too many arguments")
	}

	switch args[0] {
	case "json", "JSON":
		break
	default:
		return fmt.Errorf("invalid format: %s", args[0])
	}

	return nil
}

func JSON(data string) (string, error) {
	var curStruct interface{}

	if err := json.Unmarshal([]byte(data), &curStruct); err != nil {
		return "", fmt.Errorf("cannot unmarshal JSON: %w", err)
	}

	result, err := json.MarshalIndent(curStruct, "", "  ")
	if err != nil {
		return "", fmt.Errorf("cannot format input data: %w", err)
	}

	return string(result), nil
}
