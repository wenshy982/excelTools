package osx

import (
	"testing"
)

func TestIsFileExists(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{file: "file.go"},
			want: true,
		},
		{
			name: "test2",
			args: args{file: "file_not_exists.go"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFileExists(tt.args.file); got != tt.want {
				t.Errorf("IsFileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
