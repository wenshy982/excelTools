package randx

import (
	"testing"
)

func TestUUID(t *testing.T) {
	uuid := UUID()
	if len(uuid) != 36 {
		t.Errorf("UUID() = %v, want %v", len(uuid), 36)
	}
	t.Logf("UUID() = %v", uuid)
}

func BenchmarkUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uuid := UUID()
		if len(uuid) != 36 {
			b.Errorf("UUID() = %v, want %v", len(uuid), 36)
		}
	}
}
