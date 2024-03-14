package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runCommand = func(cmd *cobra.Command, args []string) {
	fmt.Println("execute word command")
}

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "word format convert",
	Long:  "This is a command for exchange the format of word",
	Run:   runCommand,
}

func init() {

}
