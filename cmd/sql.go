package cmd

import (
	"github.com/spf13/cobra"
)

var sqlRunFunc = func(cmd *cobra.Command, args []string) {

}

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql tool",
	Long:  "This is a sql command",
	Run:   sqlRunFunc,
}

func init() {
	sqlCmd.AddCommand(sqlStructCmd)
}
