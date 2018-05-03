package combiner

import (
	"runtime"
)

type Worker struct {
	WorkStart   int
	WorkInclude int
	WorkFinish  int

	Total int64
}

func NewWorker() *Worker { return &Worker{} }

func (exe *Worker) Start() {
	simulateWork(exe.WorkStart)
}

func (exe *Worker) Include(v Argument) {
	exe.Total += v.(int64)
	simulateWork(exe.WorkInclude)
}

func (exe *Worker) Finish() {
	simulateWork(exe.WorkFinish)
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
