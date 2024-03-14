package cmd

import (
	"github.com/spf13/cobra"
)

// add flag var here

var timeCalRunFunc = func(cmd *cobra.Command, args []string) {

}
var timeCalCmd = &cobra.Command{
	Use:   "cal",
	Short: "time cal tool",
	Long:  "This is a time caculate command",
	Run:   timeCalRunFunc,
}

func init() {
	// add subcommand here
	// set flag here
}
