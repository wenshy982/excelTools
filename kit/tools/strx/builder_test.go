package strx

import (
	"fmt"
	"strings"
	"testing"
)

// 静态字符串拼接测试

var result = "a/b/c/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t/u/v/w/x/y/z"

// 封装后的 builder 函数
// 3667611               298.9 ns/op

func TestBuilder(t *testing.T) {
	type args struct {
		parts []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{parts: []string{"a", "/", "b", "/", "c", "/", "d", "/", "e", "/", "f", "/", "g", "/", "h", "/", "i", "/", "j", "/", "k", "/", "l", "/", "m", "/", "n", "/", "o", "/", "p", "/", "q", "/", "r", "/", "s", "/", "t", "/", "u", "/", "v", "/", "w", "/", "x", "/", "y", "/", "z"}},
			want: result,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Builder(tt.args.parts...); got != tt.want {
				t.Errorf("builder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Builder("a", "/", "b", "/", "c", "/", "d", "/", "e", "/", "f", "/", "g", "/", "h", "/", "i", "/", "j", "/", "k", "/", "l", "/", "m", "/", "n", "/", "o", "/", "p", "/", "q", "/", "r", "/", "s", "/", "t", "/", "u", "/", "v", "/", "w", "/", "x", "/", "y", "/", "z")
	}
}

// 未封装的 builder 函数
// 7697828               151.6 ns/op

func BuilderStringBuilder() string {
	var builder strings.Builder
	builder.WriteString("a")
	builder.WriteString("/")
	builder.WriteString("b")
	builder.WriteString("/")
	builder.WriteString("c")
	builder.WriteString("/")
	builder.WriteString("d")
	builder.WriteString("/")
	builder.WriteString("e")
	builder.WriteString("/")
	builder.WriteString("f")
	builder.WriteString("/")
	builder.WriteString("g")
	builder.WriteString("/")
	builder.WriteString("h")
	builder.WriteString("/")
	builder.WriteString("i")
	builder.WriteString("/")
	builder.WriteString("j")
	builder.WriteString("/")
	builder.WriteString("k")
	builder.WriteString("/")
	builder.WriteString("l")
	builder.WriteString("/")
	builder.WriteString("m")
	builder.WriteString("/")
	builder.WriteString("n")
	builder.WriteString("/")
	builder.WriteString("o")
	builder.WriteString("/")
	builder.WriteString("p")
	builder.WriteString("/")
	builder.WriteString("q")
	builder.WriteString("/")
	builder.WriteString("r")
	builder.WriteString("/")
	builder.WriteString("s")
	builder.WriteString("/")
	builder.WriteString("t")
	builder.WriteString("/")
	builder.WriteString("u")
	builder.WriteString("/")
	builder.WriteString("v")
	builder.WriteString("/")
	builder.WriteString("w")
	builder.WriteString("/")
	builder.WriteString("x")
	builder.WriteString("/")
	builder.WriteString("y")
	builder.WriteString("/")
	builder.WriteString("z")
	return builder.String()
}

func TestBuilderStringBuilder(t *testing.T) {
	if got := BuilderStringBuilder(); got != result {
		t.Errorf("BuilderStringBuilder() = %v, want %v", got, result)
	}
}

func BenchmarkBuilderStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuilderStringBuilder()
	}
}

// 使用 fmt.Sprint 拼接字符串
// 1985668               620.1 ns/op

func BuilderFmt() string {
	return fmt.Sprint("a", "/", "b", "/", "c", "/", "d", "/", "e", "/", "f", "/", "g", "/", "h", "/", "i", "/", "j", "/", "k", "/", "l", "/", "m", "/", "n", "/", "o", "/", "p", "/", "q", "/", "r", "/", "s", "/", "t", "/", "u", "/", "v", "/", "w", "/", "x", "/", "y", "/", "z")
}

func TestBuilderFmt(t *testing.T) {
	if got := BuilderFmt(); got != result {
		t.Errorf("BuilderFmt() = %v, want %v", got, result)
	}
}

func BenchmarkBuilderFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuilderFmt()
	}
}

// 使用 strings.Join 拼接字符串
// 6170486               189.1 ns/op

func BuilderJoin() string {
	return strings.Join(
		[]string{
			"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o",
			"p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		},
		"/",
	)
}

func TestBuilderJoin(t *testing.T) {
	if got := BuilderJoin(); got != result {
		t.Errorf("BuilderJoin() = %v, want %v", got, result)
	}
}

func BenchmarkBuilderJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuilderJoin()
	}
}

// 使用 + 拼接字符串
// 1000000000               0.2429 ns/op

func BuilderPlus() string {
	return "a" + "/" + "b" + "/" + "c" + "/" + "d" + "/" + "e" + "/" + "f" + "/" + "g" + "/" + "h" + "/" + "i" + "/" + "j" + "/" + "k" + "/" + "l" + "/" + "m" + "/" + "n" + "/" + "o" + "/" + "p" + "/" + "q" + "/" + "r" + "/" + "s" + "/" + "t" + "/" + "u" + "/" + "v" + "/" + "w" + "/" + "x" + "/" + "y" + "/" + "z"
}

func TestBuilderPlus(t *testing.T) {
	if got := BuilderPlus(); got != result {
		t.Errorf("BuilderPlus() = %v, want %v", got, result)
	}
}

func BenchmarkBuilderPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuilderPlus()
	}
}
