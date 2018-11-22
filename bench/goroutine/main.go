package main

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/loov/hrtime"
)

func main() {
	const P = 16
	const K = 100
	const N = 10000

	{
		bench := hrtime.NewBenchmarkTSC(K)
		for bench.Next() {
			var wg sync.WaitGroup
			wg.Add(P)
			for k := 0; k < P; k++ {
				go func() {
					for i := 0; i < N; i++ {
						runtime.Gosched()
					}
					wg.Done()
				}()
			}
			wg.Wait()
		}
		hist := bench.Histogram(10)
		hist.Divide(N * P / runtime.NumCPU())
		fmt.Println("gosched\n", hist.StringStats())
	}

}
