package strx

import (
	"strings"
)

/*
Builder 拼接字符串
remark:

	凭借静态字符串时，使用 + 拼接字符串，效率最高！
	使用此方法拼接字符串，使用起来更加方便，增强代码可读性和可维护性
	但是效率比直接使用 strings.Builder.WriteString() 拼接字符串要低，大量拼接字符串时不建议使用此方法
*/
func Builder(parts ...string) string {
	var b strings.Builder
	for _, p := range parts {
		b.WriteString(p)
	}
	return b.String()
}
