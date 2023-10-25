package main

import (
	"fmt"
)

var (
	_            = Const(&exampleTable)
	exampleTable = []int{
		1, 2, 3, 4, 5, 6,
	}
)

func main() {
	fmt.Println(exampleTable)
	exampleTable[2] = 123
	fmt.Println(exampleTable[2])
}
