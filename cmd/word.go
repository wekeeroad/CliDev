package cmd

import (
	"fmt"
	"log"
	"strings"

	"clidev/pkg/word"

	"github.com/spf13/cobra"
)

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamelCase
	ModeUnderscoreToLowerCamelCase
	ModeCamelCaseToUnderscore
)

var str string
var mode int8

var wordRunFunc = func(cmd *cobra.Command, args []string) {
	var content string
	switch mode {
	case ModeUpper:
		content = word.ToUpper(str)
	case ModeLower:
		content = word.ToLower(str)
	case ModeUnderscoreToUpperCamelCase:
		content = word.UnderscoreToUpperCamelCase(str)
	case ModeUnderscoreToLowerCamelCase:
		content = word.UnderscoreToLowerCamelCase(str)
	case ModeCamelCaseToUnderscore:
		content = word.CamelCaseToUnderscore(str)
	default:
		log.Fatalf("Not support the mode, please read the usage of command by 'help word'")
	}
	fmt.Printf("Output : %s", content)
}
var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换,模式如下:",
	"1:全部单词转为大写",
	"2:全部单词转为小写",
	"3:下画线单词转为大写驼峰单词",
	"4:下画线单词转为小写驼峰单词",
	"5:驼峰单词转为下画线单词",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "word format convert",
	Long:  desc,
	Run:   wordRunFunc,
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "please input the word")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "please define the mode of word convert")
}
