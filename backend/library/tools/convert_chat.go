package tools

import "unicode"

func ConvertToLowerSnakeCase(input string) string {
	var result []rune
	for i, r := range input {
		// 如果当前字符是大写字母并且不是第一个字符，就在前面加上'.'
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, '.')
			}
			// 把大写字母转换成小写
			r = unicode.ToLower(r)
		}
		result = append(result, r)
	}
	return string(result)
}
