package combiner

import (
	"runtime"
	"time"
)

type Worker struct {
	WorkStart   int
	WorkInclude int
	WorkFinish  int

	SleepStart   time.Duration
	SleepInclude time.Duration
	SleepFinish  time.Duration

	Total int64
}

func NewWorker() *Worker { return &Worker{} }

func (exe *Worker) Start() {
	simulateWork(exe.WorkStart, exe.SleepStart)
}

func (exe *Worker) Include(v Argument) {
	exe.Total += v.(int64)
	simulateWork(exe.WorkInclude, exe.SleepInclude)
}

func (exe *Worker) Finish() {
	simulateWork(exe.WorkFinish, exe.SleepFinish)
}

func simulateWork(amount int, sleep time.Duration) {
	foo := 1
	for i := 0; i < amount; i++ {
		foo *= 2
		foo /= 2
	}
	if amount > 0 {
		runtime.Gosched()
	}

	if sleep > 0 {
		time.Sleep(sleep)
	}
}
