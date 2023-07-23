package waitgroup

import (
	"sync"
)

type WG struct {
	*sync.WaitGroup
}

// New 新建 WaitGroup
func New() *WG {
	return &WG{&sync.WaitGroup{}}
}

// Add 增加
func (wg *WG) Add(delta int) {
	wg.Add(delta)
}

// Wait 等待
func (wg *WG) Wait() {
	wg.Wait()
}

// Done 完成
func (wg *WG) Done() {
	wg.Done()
}
