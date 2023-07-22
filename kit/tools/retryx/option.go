package retryx

import (
	"time"
)

// RetryOptions 重试配置项
type RetryOptions struct {
	Attempts uint          // 重试次数
	Delay    time.Duration // 重试间隔时间
}

type RetryOption func(*RetryOptions)

func WithAttempts(attempts uint) RetryOption {
	return func(options *RetryOptions) {
		options.Attempts = attempts
	}
}

func WithDelay(delay time.Duration) RetryOption {
	return func(options *RetryOptions) {
		options.Delay = delay
	}
}

var DefaultRetryOptions = []RetryOption{
	WithAttempts(3),            // 重试3次
	WithDelay(1 * time.Second), // 每次间隔1秒
}
