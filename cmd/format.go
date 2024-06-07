/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/MircoT/go-dev-toys/pkg/format"
	"github.com/spf13/cobra"
)

// formatCmd represents the encode command
var formatCmd = &cobra.Command{
	Use:   "format targeFormat input",
	Short: "Format a string with a specific type",
	Long: `Format a string with a specific type.
	
Available formats: base64, html, url`,
	Args: format.CheckArgs,
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
		case "json", "JSON":
			result, err = format.JSON(inputString)
		}

		if err != nil {
			return fmt.Errorf(
				"cannot format input '%s' with '%s' type: %w",
				inputString, targetFormat, err,
			)
		}

		fmt.Println(result)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(formatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// formatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// formatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
