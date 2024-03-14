package cmd

import (
	"github.com/spf13/cobra"
)

// add flag var here

var timeRunFunc = func(cmd *cobra.Command, args []string) {

}
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "time tool",
	Long:  "This is a time command",
	Run:   timeRunFunc,
}

func init() {
	// add subcommand here
	// set flag here
	timeCmd.AddCommand(timeNowCmd)
	timeCmd.AddCommand(timeCalCmd)
}
