package mathx

import (
	"math"

	"github.com/shopspring/decimal"
)

// FloatAdd 浮点数相加
func FloatAdd(a, b float64) float64 {
	return DecimalRound(decimal.NewFromFloat(a).Add(decimal.NewFromFloat(b)))
}

// FloatSub 浮点数相减
func FloatSub(a, b float64) float64 {
	return DecimalRound(decimal.NewFromFloat(a).Sub(decimal.NewFromFloat(b)))
}

// FloatMul 浮点数相乘
func FloatMul(a, b float64) float64 {
	return DecimalRound(decimal.NewFromFloat(a).Mul(decimal.NewFromFloat(b)))
}

// FloatDiv 浮点数相除
func FloatDiv(a, b float64) float64 {
	return DecimalRound(decimal.NewFromFloat(a).Div(decimal.NewFromFloat(b)))
}

// FloatRound 保留小数点后两位
func FloatRound(value float64) float64 {
	return math.Round(value*100) / 100
}

// DecimalRound 保留2位小数
func DecimalRound(f decimal.Decimal) float64 {
	v, _ := f.Float64()
	return math.Round(v*100) / 100
}

// FenToPoints 分转换成积分 points = money * ratio / 100
func FenToPoints(money int64, ratio int32) float64 {
	return DecimalRound(decimal.NewFromInt(money).Mul(decimal.NewFromInt32(ratio)).Div(decimal.NewFromInt(100)))
}

// YuanToPoints 元转换成积分 points = money * ratio
func YuanToPoints(money float64, ratio int32) float64 {
	return DecimalRound(decimal.NewFromFloat(money).Mul(decimal.NewFromInt32(ratio)))
}

// PointsToFen 积分转换成分 fen = points * 100 / ratio
func PointsToFen(points float64, ratio int32) int64 {
	return int64(PointsToYuan(FloatMul(points, 100.0), ratio))
}

// PointsToYuan 积分转换成元 yuan = points / ratio
func PointsToYuan(points float64, ratio int32) float64 {
	return DecimalRound(decimal.NewFromFloat(points).Div(decimal.NewFromInt32(ratio)))
}

// CharacterToPoints 字符转换成积分 points = characterCount / ratio
func CharacterToPoints(characterCount, ratio int32) float64 {
	return DecimalRound(decimal.NewFromInt32(characterCount).Div(decimal.NewFromInt32(ratio)))
}
