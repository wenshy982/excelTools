package jsonx

import (
	"io"

	"github.com/bytedance/sonic"
	"go.uber.org/zap"
)

// Marshal 结构体转json
func Marshal(v any) []byte {
	marshal, err := sonic.Marshal(v)
	if err != nil {
		zap.L().Error("结构体转json字符串失败", zap.Error(err))
	}
	return marshal
}

// MarshalString 结构体转json字符串
func MarshalString(v any) string {
	marshal, err := sonic.MarshalString(v)
	if err != nil {
		zap.L().Error("结构体转json字符串失败", zap.Error(err))
	}
	return marshal
}

// Unmarshal json转结构体
func Unmarshal(b []byte, v any) {
	err := sonic.Unmarshal(b, v)
	if err != nil {
		zap.L().Error("json字符串转结构体失败", zap.Error(err))
	}
}

// UnmarshalString json字符串转结构体
func UnmarshalString(s string, v any) {
	err := sonic.UnmarshalString(s, v)
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
