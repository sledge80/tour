package word

import (
	"github.com/spf13/cobra"
	"log"
	"strings"
	"unicode"
)

var WordCmd *cobra.Command

func init() {
	var str string
	var mode int8

	var use = "word"
	var short = "单词格式转换"
	var desc = strings.Join([]string{
		"该子命令支持各种单词格式转换，模式如下：",
		"1：全部转大写",
		"2：全部转小写",
		"3：下划线转大写驼峰",
		"4：下划线转小写驼峰",
		"5：驼峰转下划线",
	}, "\n")
	var run = func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = ToUpper(str)
		case ModeLower:
			content = ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行 help word 查看帮助文档")
		}

		log.Printf("输出结果: %s", content)
	}

	WordCmd = &cobra.Command{
		Use:   use,
		Short: short,
		Long:  desc,
		Run:   run,
	}

	WordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	WordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}

// -------------------Mode常量定义-------------------------------------------
const (
	ModeUpper                      = iota + 1 // 全部转大写
	ModeLower                                 // 全部转小写
	ModeUnderscoreToUpperCamelCase            // 下划线转大写驼峰
	ModeUnderscoreToLowerCamelCase            // 下线线转小写驼峰
	ModeCamelCaseToUnderscore                 // 驼峰转下划线
)

// ----------------------单词格式转换工具方法------------------------------

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
