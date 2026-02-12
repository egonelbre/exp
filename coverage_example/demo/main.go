package main

import (
	sigma "ext"
	"ext/stat"
	"ext/sub"
	"fmt"
	"hello.test/cpu"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(cpu.Add(1, 2))
	fmt.Println(stat.MinMax([]int{1, 2, 3, 4, 5}))
	fmt.Println(sigma.Sum([]int{1, 2, 3, 4, 5}))
	fmt.Println(sub.Sub(1, 2))
}
