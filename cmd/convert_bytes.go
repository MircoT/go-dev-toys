/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/MircoT/go-dev-toys/pkg/convert"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

var (
	bytesUnit string
)

// bytesCmd represents the bytes command
var bytesCmd = &cobra.Command{
	Use:   "bytes NUM",
	Short: "Convert a bytes in different bases",
	Long:  `Convert a bytes in different bases.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("you have to provide a bytes to convert")
		}

		bytes, args := args[0], args[1:]

		bytesUint, err := strconv.ParseUint(bytes, 10, 0)
		if err != nil {
			return fmt.Errorf("cannot convert bytes to unsigned integer: %w", err)
		}

		res, err := convert.Bytes(bytesUint, bytesUnit)
		if err != nil {
			return fmt.Errorf("cannot convert bytes '%s': %w", bytes, err)
		}

		spew.Dump(res)
		// data := pterm.TableData{
		// 	{"Base", "Value"},
		// 	{"10", fmt.Sprintf("%d", intNum)},
		// 	{"16", fmt.Sprintf("%x", intNum)},
		// 	{"8", fmt.Sprintf("%o", intNum)},
		// 	{"2", fmt.Sprintf("%b", intNum)},
		// }

		// pterm.DefaultTable.
		// 	WithHasHeader().
		// 	WithRowSeparator("-").
		// 	WithHeaderRowSeparator("-").
		// 	WithRightAlignment().
		// 	WithData(data).
		// 	Render()

		return nil
	},
}

func init() {
	convertCmd.AddCommand(bytesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bytesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bytesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	bytesCmd.Flags().StringVar(&bytesUnit, "unit", "KiB", "Base unit of the source number")
}
