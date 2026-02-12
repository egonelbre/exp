package main

import (
	"fmt"
	"hello.test/cpu"
	"ext/stat"
	sigma "ext"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(cpu.Add(1, 2))
	fmt.Println(stat.MinMax([]int{1, 2, 3, 4, 5}))
	fmt.Println(sigma.Sum([]int{1, 2, 3, 4, 5}))
}
