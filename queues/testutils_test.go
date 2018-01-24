package goqueuestest

import (
	"sync"
	"testing"
	"time"
)

const (
	testItemCount = 1000
	testTimeout   = 10 * time.Second
)

func timeout() chan int {
	done := make(chan int)
	go func() {
		select {
		case <-time.After(testTimeout):
			panic("Timeout")
		case <-done:
		}
	}()
	return done
}

func testReturn(t *testing.T, got interface{}, wasOk bool, expected interface{}, expectedOk bool) {
	if wasOk != expectedOk {
		t.Fatalf("Returned ok=%b expected %b", wasOk, expectedOk)
	}

	if got != expected {
		t.Fatalf("Returned %v expected %v", got, expected)
	}
}

func fifoTest(t *testing.T, q Queue) {
	done := timeout()

	N := testItemCount
	for i := 0; i < N; i += 1 {
		q.Enqueue(i)
	}

	for i := 0; i < N; i += 1 {
		val, ok := q.Dequeue()
		testReturn(t, val, ok, i, true)
	}

	for i := 0; i < 5; i += 1 {
		val, ok := q.Dequeue()
		testReturn(t, val, ok, nil, false)
	}

	done <- 1
}

func lifoTest(t *testing.T, q Queue) {
	done := timeout()

	N := testItemCount
	for i := 0; i < N; i += 1 {
		q.Enqueue(i)
	}

	for i := N - 1; i >= 0; i -= 1 {
		val, ok := q.Dequeue()
		testReturn(t, val, ok, i, true)
	}

	for i := 0; i < 5; i += 1 {
		val, ok := q.Dequeue()
		testReturn(t, val, ok, nil, false)
	}
	done <- 1
}

func runParallelTest(t *testing.T, test func(int)) {
	done := timeout()
	cpus := 4
	N := testItemCount / cpus
	wg := &sync.WaitGroup{}
	for i := 0; i < cpus; i += 1 {
		wg.Add(1)
		go func(){
			test(N)
			wg.Done()
		}()
	}
	wg.Wait()
	done <- 1
}

func queuePTest(t *testing.T, q Queue) {
	runParallelTest(t, func(N int){
		for i := 0; i < N; i += 1 {
			q.Enqueue(i)
		}

		for i := 0; i < N; i += 1 {
			_, ok := q.Dequeue()
			if !ok {
				t.Fatalf("Was unable to get value")
			}
		}
	})
}

func queueP2Test(t *testing.T, q Queue) {
	runParallelTest(t, func(N int){
		for i := 0; i < N; i += 1 {
			q.Enqueue(i)
			_, ok := q.Dequeue()
			if !ok {
				t.Fatalf("Was unable to get value")
			}
		}
	})
}
