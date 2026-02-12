package main

import (
	"testing"

	sigma "ext"
	"ext/stat"
	"ext/sub"
	"hello.test/cpu"
)

func TestMain(t *testing.T) {
	t.Log("Hello, World!")
	t.Log(cpu.Add(1, 2))
	t.Log(stat.MinMax([]int{1, 2, 3, 4, 5}))
	t.Log(sigma.Sum([]int{1, 2, 3, 4, 5}))
	t.Log(sub.Sub(1, 2))
}
