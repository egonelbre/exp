package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync/atomic"
	"time"

	"github.com/egonelbre/exp/ring"
)

var (
	TotalWritten int64
	TotalRead    int64
)

func writer(data *ring.Ring) {
	for i := 0; i < 1<<20; i++ {
		cell := data.Allocate()
		cell.Size = rand.Int63() & 0xFF
		atomic.AddInt64(&TotalWritten, cell.Size)
		cell.Enqueue()

		runtime.Gosched()
	}
}

func reader(data *ring.Ring) {
	for {
		cell := data.Dequeue()
		if cell == nil {
			panic("invalid value")
		}

		atomic.AddInt64(&TotalRead, cell.Size)
		cell.Release()
	}
}

func monitor(data *ring.Ring) {
	for {
		written, read := atomic.LoadInt64(&TotalWritten), atomic.LoadInt64(&TotalRead)
		fmt.Printf("%d/%d\n", written, read)
		time.Sleep(time.Second)
	}
}

func main() {
	data := ring.New()
	go monitor(data)
	for i := 0; i < 64; i++ {
		go reader(data)
	}
	go writer(data)

	time.Sleep(10 * time.Second)
}
