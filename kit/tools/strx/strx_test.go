package strx

import (
	"testing"
)

func TestCountCharacters(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "test1",
			args: args{message: "hello"},
			want: 5,
		},
		{
			name: "test2",
			args: args{message: "hello world"},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountCharacters(tt.args.message); got != tt.want {
				t.Errorf("CountCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCountCharacters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountCharacters("hello world")
	}
}

func TestMaskPhoneNumber(t *testing.T) {
	type args struct {
		phoneNumber string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{phoneNumber: "18250300036"},
			want: "182****0036",
		},
		{
			name: "test2",
			args: args{phoneNumber: "182503"},
			want: "182503",
		},
		{
			name: "test3",
			args: args{phoneNumber: ""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaskPhoneNumber(tt.args.phoneNumber); got != tt.want {
				t.Errorf("MaskPhoneNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMaskPhoneNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaskPhoneNumber("18250300036")
	}
}

func TestTruncate(t *testing.T) {
	type args struct {
		s         string
		maxLength int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Truncate",
			args: args{
				s:         "1234567890",
				maxLength: 10,
			},
			want: "1234567890",
		},
		{
			name: "Truncate",
			args: args{
				s:         "hello world",
				maxLength: 5,
			},
			want: "hello...",
		},
		{
			name: "test1",
			args: args{
				s:         "This is a test string.",
				maxLength: 3,
			},
			want: "Thi...",
		},
		{
			name: "test2",
			args: args{s: "中文测试字符串",
				maxLength: 3,
			},
			want: "中文测...",
		},
		{
			name: "test3",
			args: args{s: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
				maxLength: 3,
			},
			want: "xxx...",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Truncate(tt.args.s, tt.args.maxLength); got != tt.want {
				t.Errorf("Truncate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTruncate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Truncate("This is a test string.", 3)
	}
}

func TestAddBackslashBeforeQuotes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{s: "This is a \"quote\" and this is a backquote: `.`"},
			want: "This is a \\\"quote\\\" and this is a backquote: \\`.\\`",
		},
		{
			name: "test2",
			args: args{s: "Hello \"world\""},
			want: "Hello \\\"world\\\"",
		},
		{
			name: "test3",
			args: args{s: "This is a \"test\" with multiple \"quotes\" and `backquotes`."},
			want: "This is a \\\"test\\\" with multiple \\\"quotes\\\" and \\`backquotes\\`.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddBackslashBeforeQuotes(tt.args.s); got != tt.want {
				t.Errorf("AddBackslashBeforeQuotes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAddBackslashBeforeQuotes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddBackslashBeforeQuotes("This is a \"quote\" and this is a backquote: `.`")
	}
}

func TestReplaceBackslashes(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{path: "C:\\Users\\Administrator\\Desktop\\test.txt"},
			want: "C:/Users/Administrator/Desktop/test.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceBackslashes(tt.args.path); got != tt.want {
				t.Errorf("ReplaceBackslashes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkReplaceBackslashes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReplaceBackslashes("C:\\Users\\Administrator\\Desktop\\test.txt")
	}
}
