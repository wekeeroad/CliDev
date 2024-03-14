package cmd

import (
	"clidev/pkg/timer"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// add flag var here
var calculateTime string
var duration string

var timeCalRunFunc = func(cmd *cobra.Command, args []string) {
	var location, _ = time.LoadLocation("Asia/Shanghai")
	var currentTime time.Time
	var layout = "2006-01-02 15:04:05"
	if calculateTime == "" {
		currentTime = timer.GetNowTime()
	} else {
		var err error
		space := strings.Count(calculateTime, " ")
		if space == 0 {
			layout = "2006-01-02"
		} else {
			div := strings.Count(calculateTime, ":")
			switch div {
			case 0:
				layout = "2006-01-02 15"
			case 1:
				layout = "2006-01-02 15:04"
			case 2:
				layout = "2006-01-02 15:04:05"
			}
		}
		currentTime, err = time.ParseInLocation(layout, calculateTime, location)
		if err != nil {
			t, _ := strconv.Atoi(calculateTime)
			currentTime = time.Unix(int64(t), 0)
		}
	}
	t, err := timer.GetCaculateTime(currentTime, duration)
	if err != nil {
		log.Fatalf("timer.GetCaculateTime err:%v", err)
	}
	fmt.Printf("Output: %s, %d", t.In(location).Format(layout), t.Unix())
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
	timeCalCmd.Flags().StringVarP(&calculateTime, "alculate", "c", "", "The start of calculate time")
	timeCalCmd.Flags().StringVarP(&duration, "duration", "d", "", "The increase of time for calculation")
}
