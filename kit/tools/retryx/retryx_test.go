package retryx

import (
	"errors"
	"testing"
)

func TestDo(t *testing.T) {
	type args struct {
		retryFunction func() error
		opts          []RetryOption
	}

	attempts := 0

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test success",
			args: args{
				retryFunction: func() error { return nil },
				opts: []RetryOption{
					WithAttempts(3),
					WithDelay(1000),
				},
			},
			wantErr: false,
		},
		{
			name: "test fail",
			args: args{
				retryFunction: func() error {
					return errors.New("test error")
				},
				opts: []RetryOption{
					WithAttempts(3),
					WithDelay(1000),
				},
			},
			wantErr: true,
		},
		{
			name: "test mid success",
			args: args{
				retryFunction: func() error {
					attempts++
					if attempts <= 2 {
						return errors.New("test error")
					}
					return nil
				},
				opts: []RetryOption{
					WithAttempts(3),
					WithDelay(1000),
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Do(tt.args.retryFunction, tt.args.opts...); (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
