package mathx

import (
	"sync"
	"testing"
)

func TestAtomicAddInt32(t *testing.T) {
	var v int32 = 0
	var times = 1000000
	var n = 3
	var wg sync.WaitGroup
	wg.Add(n)
	go func() {
		defer wg.Done()
		for i := 0; i < times; i++ {
			AtomicAddInt32(&v, 1)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < times; i++ {
			AtomicAddInt32(&v, 1)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < times; i++ {
			AtomicAddInt32(&v, 1)
		}
	}()
	wg.Wait()
	if v != int32(times*n) {
		t.Errorf("AtomicSubInt32() = %v, want %v", v, 0)
	}
	t.Logf("v = %v", v)
}

func BenchmarkAtomicAddInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AtomicAddInt32(&[]int32{1, 2, 3}[i%3], 1)
	}
}

func TestAtomicSubInt32(t *testing.T) {
	var v int32 = 3000000
	var times int32 = 1000000
	var wg sync.WaitGroup
	wg.Add(int(v / times))
	go func() {
		defer wg.Done()
		for i := int32(0); i < times; i++ {
			AtomicSubInt32(&v, 1)
		}
	}()
	go func() {
		defer wg.Done()
		for i := int32(0); i < times; i++ {
			AtomicSubInt32(&v, 1)
		}
	}()
	go func() {
		defer wg.Done()
		for i := int32(0); i < times; i++ {
			AtomicSubInt32(&v, 1)
		}
	}()
	wg.Wait()
	if v != 0 {
		t.Errorf("AtomicSubInt32() = %v, want %v", v, 0)
	}
	t.Logf("v = %v", v)
}

func BenchmarkAtomicSubInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AtomicSubInt32(&[]int32{1, 2, 3}[0], 1)
	}
}
