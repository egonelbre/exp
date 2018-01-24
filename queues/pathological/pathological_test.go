package pathological
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

// ScLifo
func BenchmarkScLifo_ANRNx2(b *testing.B) { runParallelBench(b, 2, NewScLifo(), benchANRN)}
func BenchmarkScLifo_ANRNx3(b *testing.B) { runParallelBench(b, 3, NewScLifo(), benchANRN)}

// SmLifo
func BenchmarkSmLifo_ANRNx2(b *testing.B) { runParallelBench(b, 2, NewSmLifo(), benchANRN)}
func BenchmarkSmLifo_ANRNx3(b *testing.B) { runParallelBench(b, 3, NewSmLifo(), benchANRN)}
