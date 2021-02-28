package cmd

import (
	"NanjoFan/Tour/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

const (
	MODE_UPPER                         = iota + 1 // 全部单词转为大写
	MODE_LOWER                                    // 全部转为小写
	MODE_UNDERSCORE_TO_UPPER_CAMELCASE            // 下划线->大驼峰
	MODE_UNDERSCORE_TO_LOWER_CAMELCASE            // 下划线->小驼峰
	MODE_CAMELCASE_TO_UNDERSCORE                  // 驼峰单词转为下划线单词
)

var str string
var mode int8

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转化，模式如下：",
	"1：全部转为大写",
	"2：全部转为小写",
	"3：下划线->大驼峰",
	"4：下划线->小驼峰",
	"5：驼峰单词转为下划线单词",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case MODE_UPPER:
			content = word.ToUpper(str)
		case MODE_LOWER:
			content = word.ToLower(str)
		case MODE_UNDERSCORE_TO_UPPER_CAMELCASE:
			content = word.UnderscoreToUpperCamelCase(str)
		case MODE_UNDERSCORE_TO_LOWER_CAMELCASE:
			content = word.UnderscoreToLowerCamelCase(str)
		case MODE_CAMELCASE_TO_UNDERSCORE:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalln("暂不支持次格式，请执行 help word 查看帮助支持")
		}
		log.Printf("输出结果： %s\n", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}
