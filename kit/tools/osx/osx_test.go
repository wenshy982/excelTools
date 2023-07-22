package osx

import (
	"os"
	"testing"
	"time"
)

var wd, _ = os.Getwd()
var dirname = "test"

func TestPwd(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test1",
			want: wd,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pwd(); got != tt.want {
				t.Errorf("Pwd() = %v, want %v", got, tt.want)
			}
		})
	}
	t.Logf("pwd: %s", Pwd())
}

func TestDirParent(t *testing.T) {
	_ = MkdirAll(dirname)
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{dir: dirname},
			want: ".", // current dir
		},
		// {
		// 	name: "test2",
		// 	args: args{dir: wd},
		// 	want: `F:\desktop\excelTool\kit\tools`,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DirParent(tt.args.dir); got != tt.want {
				t.Errorf("DirParent() = %v, want %v", got, tt.want)
			}
		})
	}
	_ = RemoveAll(dirname)
}

func TestPwdParent(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test1",
			want: DirParent(wd),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PwdParent(); got != tt.want {
				t.Errorf("PwdParent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMkdirAll(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				dir: dirname,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MkdirAll(tt.args.dir); (err != nil) != tt.wantErr {
				t.Errorf("Mkdir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	_ = RemoveAll(dirname)
}

func TestRemoveAll(t *testing.T) {
	_ = MkdirAll(dirname)
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test1",
			args:    args{path: dirname},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RemoveAll(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("RemoveAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRemoveTimeoutDir(t *testing.T) {
	_ = MkdirAll(dirname)
	type args struct {
		dir     string
		timeout time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test1",
			args:    args{dir: dirname, timeout: time.Second},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(1 * time.Second)
			err := RemoveTimeoutDir(tt.args.dir, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveTimeoutDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsPathExist(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				dir: dirname,
			},
			want: false,
		},
		{
			name: "test1",
			args: args{
				dir: DirParent(dirname),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPathExist(tt.args.dir); got != tt.want {
				t.Errorf("IsPathExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitDir(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{dir: dirname},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitDir(tt.args.dir)
		})
	}
	_ = RemoveAll(dirname)
}
