package tokenx

import (
	"testing"
)

func TestCalToken(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{msg: "你好"},
			want: 2,
		},
		{
			name: "test2",
			args: args{msg: "Hello World"},
			want: 2,
		},
		{
			name: "test3",
			args: args{msg: "Hello World, 你好世界"},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalToken(tt.args.msg); got != tt.want {
				t.Errorf("CalToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
