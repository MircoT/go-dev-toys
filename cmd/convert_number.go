/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var (
	numberBase int
)

// numberCmd represents the number command
var numberCmd = &cobra.Command{
	Use:   "number NUM",
	Short: "Convert a number in different bases",
	Long:  `Convert a number in different bases.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if numberBase != 10 && numberBase != 2 && numberBase != 8 && numberBase != 16 {
			return fmt.Errorf("base %d is not supported", numberBase)
		}

		if len(args) != 1 {
			return fmt.Errorf("you have to provide a number to convert")
		}

		number, args := args[0], args[1:]

		intNum, err := strconv.ParseInt(number, numberBase, 0)
		if err != nil {
			return fmt.Errorf("cannot parse int number '%s': %w", number, err)
		}

		data := pterm.TableData{
			{"Base", "Value"},
			{"10", fmt.Sprintf("%d", intNum)},
			{"16", fmt.Sprintf("%x", intNum)},
			{"8", fmt.Sprintf("%o", intNum)},
			{"2", fmt.Sprintf("%b", intNum)},
		}

		pterm.DefaultTable.
			WithHasHeader().
			WithRowSeparator("-").
			WithHeaderRowSeparator("-").
			WithRightAlignment().
			WithData(data).
			Render()

		return nil
	},
}

func init() {
	convertCmd.AddCommand(numberCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// numberCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// numberCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	numberCmd.Flags().IntVar(&numberBase, "base", 10, "Base of the source number")
}
