package timex

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"time"
)

// RandomMillisecond 获取指定范围内的随机毫秒数
func RandomMillisecond(n int64) time.Duration {
	if n < 50 {
		n = 50
	}
	rnd, _ := rand.Int(rand.Reader, big.NewInt(n))
	return time.Duration(int(rnd.Int64())) * time.Millisecond
}

// Timeout 检查是否超时 true: 已经超时 false: 没有超时
func Timeout(checkTime time.Time, timeout time.Duration) bool {
	return checkTime.Add(timeout).Before(time.Now())
}

// Cost 计算耗时
func Cost() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		log.Printf("【处理完成】耗时：%v 按回车键退出... \n", tc)
		_, _ = fmt.Scanln()
	}
}

// ResetTimer 重置定时器
func ResetTimer(timer *time.Ticker, duration time.Duration) {
	if timer != nil {
		timer.Stop() // 停止旧的定时器
	}
	timer = time.NewTicker(duration)
}
