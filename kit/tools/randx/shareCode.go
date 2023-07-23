package randx

import (
	cRand "crypto/rand"
	"io"
	"strings"
)

const (
	shareCodeLen     = 5                                                                // 分享码长度
	shareCodeCharSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // 分享码字符集
)

// ShareCode 生成分享码
func ShareCode() string {
	b := make([]byte, shareCodeLen)
	_, _ = io.ReadFull(cRand.Reader, b)
	for i := 0; i < shareCodeLen; i++ {
		b[i] = shareCodeCharSet[int(b[i])%len(shareCodeCharSet)]
	}
	var builder strings.Builder
	builder.Write(b)
	return builder.String()
}
