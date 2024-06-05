/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/MircoT/go-dev-toys/pkg/dectext/zstd"
	"github.com/spf13/cobra"
)

// decompressCmd represents the decompress command
var decompressCmd = &cobra.Command{
	Use:   "decompress file",
	Short: "Decompress a compressed text",
	Long:  `Decompress a compressed text.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			inputString []byte
			result      string
			err         error
		)

		targetFile, args := args[0], args[1:]

		inputString, err = os.ReadFile(targetFile)
		if err != nil {
			return fmt.Errorf("cannot read file %s: %w", targetFile, err)
		}

		result, err = zstd.Decompress(inputString)
		if err != nil {
			return fmt.Errorf("cannot decompress text: %w", err)
		}

		fmt.Println(result)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(decompressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decompressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decompressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
