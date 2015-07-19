package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/egonelbre/exp/aq"
)

const MaxSize = 512 << 20

func main() {
	db, err := aq.New("example.db", MaxSize)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	const (
		minWrite = 1 << 8
		maxWrite = 1 << 12
	)

	empty := make([]byte, maxWrite)
	written := 0
	x := 0

	start := time.Now()
	for {
		n := rand.Intn(maxWrite-minWrite) + minWrite
		_, err := db.Write(empty[:n])
		if err != nil {
			break
		}

		written += n
		x++
	}
	db.Flush()
	elapsed := time.Since(start)

	mb := float64(written) / float64(1<<20)
	fmt.Printf("Write %.3fMB in %v\n", mb, elapsed)
	fmt.Printf("      %.3fMB/s\n", mb/elapsed.Seconds())
	fmt.Printf("      %v entries\n", x)

	start = time.Now()
	read := 0
	x = 0
	it := db.Iterate()
	for it.Next() {
		n := len(it.Bytes())
		read += n
		x++
	}
	elapsed = time.Since(start)

	mb = float64(read) / float64(1<<20)
	fmt.Printf("Read  %.3fMB in %v\n", mb, elapsed)
	fmt.Printf("      %.3fMB/s\n", mb/elapsed.Seconds())
	fmt.Printf("      %v entries\n", x)
}
