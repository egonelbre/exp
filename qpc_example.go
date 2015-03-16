package main

import (
	"math/rand"
	"time"

	"./qpc"
)

func main() {
	A := qpc.NewHistory("A")
	B := qpc.NewHistory("B")

	for i := 0; i < 1000; i += 1 {
		A.Start()
		time.Sleep(time.Duration(rand.Intn(1000)))
		A.Stop()

		B.Start()
		time.Sleep(time.Duration(rand.Intn(1000) + 1000))
		B.Stop()
	}

	qpc.PrintSummary(A, B)
}
