/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/MircoT/go-dev-toys/pkg/encdec"
	"github.com/MircoT/go-dev-toys/pkg/encdec/base64"
	"github.com/MircoT/go-dev-toys/pkg/encdec/html"
	"github.com/MircoT/go-dev-toys/pkg/encdec/url"
	"github.com/spf13/cobra"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode string in various formats",
	Long: `Decode string in various formats.
	
Available formats: base64, html, url`,
	Args: encdec.CheckArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			targetFormat string
			inputString  string
			result       string
			err          error
		)

		targetFormat, args = args[0], args[1:]
		inputString, args = args[0], args[1:]

		switch targetFormat {
		case "base64", "b64":
			result, err = base64.Decode(inputString)
		case "html":
			result, err = html.Decode(inputString)
		case "url", "URL":
			result, err = url.Decode(inputString)
		}

		if err != nil {
			return fmt.Errorf(
				"cannot decode input '%s' into '%s' format: %w",
				inputString, targetFormat, err,
			)
		}

		fmt.Println(result)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
