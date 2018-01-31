package main

import "fmt"

func main() {
	board := NewBoard()

	for y := 0; y < Size; y++ {
		for x := 0; x < Size; x++ {
			board.Set(x, y, RandomColor())
		}
	}

	fmt.Println(board.String())
}
