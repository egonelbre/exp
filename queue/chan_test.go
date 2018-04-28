package queue

import "testing"

func TestChan(t *testing.T) {
	test(t, func(size int) Queue { return NewChan(size) })
}
