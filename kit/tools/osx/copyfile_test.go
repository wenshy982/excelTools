package osx

import (
	"testing"
)

func TestCopyFile(t *testing.T) {
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				src: "file.go",
				dst: "file_copy.go",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CopyFile(tt.args.src, tt.args.dst)
			if !IsFileExists(tt.args.dst) {
				t.Errorf("CopyFile() error")
			}
			_ = RemoveAll(tt.args.dst)
		})
	}
}
