package goqueuestest

import (
	"sync"
	"testing"
)

const (
	useFixedBench = true
	benchIterations = 5000000
)

func runParallelBench(b *testing.B, numRoutines int, q Queue, test func(Queue, int)) {
	if useFixedBench {
		b.N = benchIterations
	}
	N := b.N / numRoutines
	wg := &sync.WaitGroup{}
	for i := 0; i < numRoutines; i += 1 {
		wg.Add(1)
		go func(){
			test(q, N)
			wg.Done()
		}()
	}
	wg.Wait()
}

func benchANRN(q Queue, N int){
	for i := 0; i < N; i += 1 {
		q.Enqueue(i)
	}
	for i := 0; i < N; i += 1 {
		q.Dequeue()
	}
}

func benchA1R1(q Queue, N int){
	for i := 0; i < N; i += 1 {
		q.Enqueue(i)
		q.Dequeue()
	}
}

func benchA1R2(q Queue, N int){
	N2 := N / 2
	for i := 0; i < N2; i += 1 {
		q.Enqueue(i)
	}
	for i := 0; i < N2; i += 1 {
		q.Enqueue(i)
		q.Dequeue()
		q.Dequeue()
	}
}

func benchA2R1(q Queue, N int){
	N2 := N / 2
	for i := 0; i < N2; i += 1 {
		q.Enqueue(i)
		q.Enqueue(i)
		q.Dequeue()
	}
	for i := 0; i < N2; i += 1 {
		q.Dequeue()
	}
}

