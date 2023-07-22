package jsonx

import (
	"io"

	"github.com/bytedance/sonic"
	"go.uber.org/zap"
)

// Marshal 结构体转json字符串
func Marshal(v any) string {
	marshal, err := sonic.Marshal(v)
	if err != nil {
		zap.L().Error("结构体转json字符串失败", zap.Error(err))
	}
	return string(marshal)
}

// Unmarshal json字符串转结构体
func Unmarshal(s string, v any) {
	err := sonic.Unmarshal([]byte(s), v)
	if err != nil {
		zap.L().Error("json字符串转结构体失败", zap.Error(err))
	}
}

// DecodeResJSON 解析http请求返回的json数据
func DecodeResJSON(body io.Reader, res any) error {
	err := sonic.ConfigDefault.NewDecoder(body).Decode(&res)
	if err != nil {
		zap.L().Error("json字符串转结构体失败", zap.Error(err))
	}
	return nil
}
