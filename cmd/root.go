package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "clidev",
	Short: "Just demo for clidev",
	Long:  "This is a demo of clidev",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
