package strx

import (
	"strconv"
)

// ToFloat64 转换为float64
func ToFloat64(value string) float64 {
	if value == "" {
		return 0
	}
	float, _ := strconv.ParseFloat(value, 64)
	return float
}

// ToInt 转换为int
func ToInt(value string) int {
	if value == "" {
		return 0
	}
	intVal, _ := strconv.Atoi(value)
	return intVal
}
