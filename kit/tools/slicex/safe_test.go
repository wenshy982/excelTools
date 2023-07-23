package slicex

import (
	"fmt"
	"sync"
	"testing"
)

func TestSafeSlice(t *testing.T) {
	s := &SafeSlice{}

	wg := sync.WaitGroup{}
	wg.Add(2)

	// Goroutine 1
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			s.Append(i)
		}
	}()

	// Goroutine 2
	go func() {
		defer wg.Done()
		for i := 1000; i < 2000; i++ {
			s.Append(i)
		}
	}()

	wg.Wait()

	// 打印完整的切片，检查是否正确
	for _, v := range s.Data {
		fmt.Println(v)
	}
}
