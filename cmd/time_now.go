package cmd

import (
	"github.com/spf13/cobra"
)

// add flag var here

var timeNowRunFunc = func(cmd *cobra.Command, args []string) {

}
var timeNowCmd = &cobra.Command{
	Use:   "now",
	Short: "time now tool",
	Long:  "This is a time now command",
	Run:   timeNowRunFunc,
}

func init() {
	// add subcommand here
	// set flag here
}
