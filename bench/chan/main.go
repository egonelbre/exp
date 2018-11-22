package main

import (
	"fmt"
	"sync"

	"github.com/loov/hrtime"
)

func main() {
	const P = 4
	const K = 100
	const N = 1000000

	{
		bench := hrtime.NewBenchmarkTSC(K)
		for bench.Next() {
			_ = make(chan int, N)
		}
		hist := bench.Histogram(10)
		hist.Divide(N)
		fmt.Println("make\n", hist.StringStats())
	}

	{
		bench := hrtime.NewBenchmarkTSC(K)
		for bench.Next() {
			ch := make(chan int, N)
			for i := 0; i < N; i++ {
				ch <- i
			}
		}
		hist := bench.Histogram(10)
		hist.Divide(N)
		fmt.Println("uncontended send\n", hist.StringStats())
	}

	{
		bench := hrtime.NewBenchmarkTSC(K)
		for bench.Next() {
			ch := make(chan int, N)
			for i := 0; i < N; i++ {
				ch <- i
			}
			for i := 0; i < N; i++ {
				<-ch
			}
		}
		hist := bench.Histogram(10)
		hist.Divide(N)
		fmt.Println("uncontended send/recv\n", hist.StringStats())
	}

	{
		bench := hrtime.NewBenchmarkTSC(K)
		for bench.Next() {
			ch := make(chan int, N)
			var wg sync.WaitGroup
			wg.Add(P)
			for k := 0; k < P; k++ {
				go func() {
					for i := 0; i < N/P; i++ {
						ch <- i
					}
					wg.Done()
				}()
			}
			wg.Wait()
		}
		hist := bench.Histogram(10)
		hist.Divide(N)
		fmt.Println("contended send\n", hist.StringStats())
	}

	{
		bench := hrtime.NewBenchmarkTSC(K)
		for bench.Next() {
			ch := make(chan int, N)
			var wg sync.WaitGroup

			wg.Add(P)
			for k := 0; k < P; k++ {
				go func() {
					for i := 0; i < N/P; i++ {
						ch <- i
					}
					wg.Done()
				}()
			}
			wg.Wait()

			wg.Add(P)
			for k := 0; k < P; k++ {
				go func() {
					for i := 0; i < N/P; i++ {
						<-ch
					}
					wg.Done()
				}()
			}
			wg.Wait()
		}
		hist := bench.Histogram(10)
		hist.Divide(N)
		fmt.Println("contended send/recv\n", hist.StringStats())
	}
}
