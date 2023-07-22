package tokenx

import (
	"github.com/tiktoken-go/tokenizer"
	"go.uber.org/zap"
)

// CalToken 计算token数量
func CalToken(msg string) int {
	codec, _ := tokenizer.Get(tokenizer.Cl100kBase)
	ids, _, _ := codec.Encode(msg)
	zap.L().Info("token数量", zap.Int("数量", len(ids)))
	return len(ids)
}
