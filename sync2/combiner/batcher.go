package combiner

import (
	"runtime"
	"sync"
)

type Worker struct {
	WorkStart   int
	WorkInclude int
	WorkFinish  int

	mu    sync.Mutex
	total int64
}

func NewWorker() *Worker { return &Worker{} }

func (exe *Worker) Start() {
	exe.mu.Lock()
	simulateWork(exe.WorkStart)
}

func (exe *Worker) Include(v Argument) {
	exe.total += v.(int64)
	simulateWork(exe.WorkInclude)
}

func (exe *Worker) Finish() {
	simulateWork(exe.WorkFinish)
	exe.mu.Unlock()
}

func simulateWork(amount int) {
	foo := 1
	for i := 0; i < amount; i++ {
		foo *= 2
		foo /= 2
	}
	if amount > 0 {
		runtime.Gosched()
	}
}
