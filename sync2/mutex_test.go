package sync2_test

import (
	"sync"
	"testing"

	"github.com/egonelbre/exp/sync2"
)

func TestMutexOwn(t *testing.T) {
	var mu sync2.Mutex
	mu.Lock()
	defer mu.Unlock()
	if !mu.Own() {
		t.Errorf("owning returned wrong result")
	}
}

func TestMutexNotOwned(t *testing.T) {
	var mu sync2.Mutex
	if mu.Own() {
		t.Errorf("should not be owning")
	}
}

func TestMutexContention(t *testing.T) {
	var wg sync.WaitGroup

	var mu sync2.Mutex
	mu.Lock()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if err := recover(); err != nil {
				return
			}
			t.Errorf("did not detect failure")
		}()

		mu.MustOwn()
	}()
	wg.Wait()
	mu.Unlock()

	if mu.Own() {
		t.Errorf("should not be owning")
	}
}
