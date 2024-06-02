package encdec

import (
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
	case "base64", "b64":
		break
	case "html":
		break
	case "URL", "url":
		break
	default:
		return fmt.Errorf("invalid format: %s", args[0])
	}

	return nil
}
