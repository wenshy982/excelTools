package ptr

import (
	"reflect"
	"testing"
	"time"
)

func TestTo(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "test1 - int",
			args: args{value: 42},
			want: 42,
		},
		{
			name: "test2 - string",
			args: args{value: "Hello, World!"},
			want: "Hello, World!",
		},
		{
			name: "test3 - float64",
			args: args{value: 42.42},
			want: 42.42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.args.value.(type) {
			case int:
				ptr := To[int](v)
				if *ptr != tt.want {
					t.Errorf("Ptr() = %v, want %v", *ptr, tt.want)
				}
			case string:
				ptr := To[string](v)
				if *ptr != tt.want {
					t.Errorf("Ptr() = %v, want %v", *ptr, tt.want)
				}
			case float64:
				ptr := To[float64](v)
				if *ptr != tt.want {
					t.Errorf("Ptr() = %v, want %v", *ptr, tt.want)
				}
			default:
				t.Errorf("Unsupported type: %T", tt.args.value)
			}
		})
	}
}

func TestNow(t *testing.T) {
	tests := []struct {
		name string
		want *time.Time
	}{
		{
			name: "test1",
			want: To(time.Now()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Now(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Now() = %v, want %v", got, tt.want)
			}
		})
	}
}
