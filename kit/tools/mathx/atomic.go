package mathx

import (
	"sync/atomic"
)

// AtomicAddInt32 原子加法
func AtomicAddInt32(addr *int32, delta int32) int32 {
	return atomic.AddInt32(addr, delta)
}

// AtomicSubInt32 原子减法
func AtomicSubInt32(addr *int32, delta int32) int32 {
	return AtomicAddInt32(addr, -delta)
}
