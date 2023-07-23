package slicex

import (
	"sync"
)

// SafeSlice 保证并发访问切片的安全
type SafeSlice struct {
	Data []any
	mu   sync.Mutex
}

func (s *SafeSlice) Append(item any) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Data = append(s.Data, item)
}

func (s *SafeSlice) Get(index int) (any, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index < 0 || index >= len(s.Data) {
		return 0, false
	}
	return s.Data[index], true
}
