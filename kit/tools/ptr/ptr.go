package ptr

import (
	"time"
)

// To 将任意类型转换为指针类型
func To[T any](t T) *T {
	return &t
}

// Now 返回当前时间的指针
func Now() *time.Time {
	return To(time.Now())
}
