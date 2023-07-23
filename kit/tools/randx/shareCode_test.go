package randx

import (
	"strings"
	"testing"
)

func TestShareCode(t *testing.T) {
	shareCode := ShareCode()
	if len(shareCode) != shareCodeLen {
		t.Errorf("ShareCode() = %v, want %v", len(shareCode), 5)
	}
	for _, c := range shareCode {
		if !strings.Contains(shareCodeCharSet, string(c)) {
			t.Errorf("ShareCode() = %v, want %v", c, shareCodeCharSet)
		}
	}
	t.Logf("ShareCode() = %v", shareCode)
}

func BenchmarkShareCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		shareCode := ShareCode()
		if len(shareCode) != shareCodeLen {
			b.Errorf("ShareCode() = %v, want %v", len(shareCode), 5)
		}
		for _, c := range shareCode {
			if !strings.Contains(shareCodeCharSet, string(c)) {
				b.Errorf("ShareCode() = %v, want %v", c, shareCodeCharSet)
			}
		}
	}
}
