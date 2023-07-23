package randx

import (
	"testing"
)

func TestInt(t *testing.T) {
	x := 6
	v := Int(x)
	if len(v) != x {
		t.Errorf("Int() = %v, want %v", len(v), x)
	}
	t.Logf("Int() = %v", v)
}

func BenchmarkInt(b *testing.B) {
	x := 6
	for i := 0; i < b.N; i++ {
		v := Int(x)
		if len(v) != x {
			b.Errorf("Int() = %v, want %v", len(v), x)
		}
	}
}
