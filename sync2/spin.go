package sync2

import "sync/atomic"

type SpinMutex struct {
	locked int64
	_      [7]int64
}

func (m *SpinMutex) Lock() {
	for atomic.SwapInt64(&m.locked, 1) == 1 {
		for atomic.LoadInt64(&m.locked) == 1 {
		}
	}
}
func (m *SpinMutex) Unlock() {
	atomic.StoreInt64(&m.locked, 0)
}
