package word

import (
	"strings"
	"unicode"
)

// 全部转为大写字母
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// 全部转为小写字母
func ToLower(s string) string {
	return strings.ToLower(s)
}

// 下划线单词转为大写驼峰单词（大驼峰式命名法（upper camel case））
// 每一个单字的首字母都采用大写字母，例如：FirstName、LastName、CamelCase，也被称为 Pascal 命名法。
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// 下划线单词转为小写驼峰单词 （小驼bai峰式命名法（lower camel case））
// 第一个单字以小写字母开始，第二个单字的首字母大写。例如：firstName、lastName。
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// 驼峰单词转换为下划线单词
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		// 对第一个字母做特殊处理
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
