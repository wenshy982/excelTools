package retryx

import (
	"github.com/avast/retry-go"
)

// Do 函数重试
// 重试机制：
//
//	重试3次
//	每次间隔1秒 (1000毫秒)
func Do(retryFunction func() error, opts ...RetryOption) error {
	var options RetryOptions

	for _, opt := range DefaultRetryOptions {
		opt(&options)
	}

	for _, opt := range opts {
		opt(&options)
	}

	err := retry.Do(
		func() error {
			err := retryFunction()
			if err != nil {
				return err
			}
			return nil
		},
		retry.Attempts(options.Attempts),
		retry.Delay(options.Delay),
	)

	if err != nil {
		return err
	}

	return nil
}
