package strx

import (
	"strings"
	"unicode/utf8"
)

// CountCharacters 统计字符串中字符个数
func CountCharacters(message string) int32 {
	return int32(utf8.RuneCount([]byte(message)))
}

// MaskPhoneNumber 隐藏手机号中间四位
func MaskPhoneNumber(phoneNumber string) string {
	if len(phoneNumber) == 11 {
		return phoneNumber[0:3] + strings.Repeat("*", 4) + phoneNumber[7:11]
	}
	return phoneNumber
}

// Truncate 截取字符串 超过maxLength字符省略
func Truncate(s string, maxLength int) string {
	r := []rune(s)
	count := 0
	for i := 0; i < len(r); {
		_, size := utf8.DecodeRuneInString(s[i:])
		count += size
		if count > maxLength {
			return string(r[:i]) + "..."
		}
		i += size
	}
	return s
}

// AddBackslashBeforeQuotes 在双引号和反引号前添加反斜杠
func AddBackslashBeforeQuotes(s string) string {
	s = strings.Replace(s, `"`, `\"`, -1)
	s = strings.Replace(s, "`", "\\`", -1)
	return s
}

// ReplaceBackslashes 将反斜杠替换为正斜杠
func ReplaceBackslashes(path string) string {
	return strings.ReplaceAll(path, "\\", "/")
}
