package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/egonelbre/exp/coder/arith"
	"github.com/egonelbre/exp/physicscompress/physics"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var (
	tries = flag.Int("tries", 1e6, "how many random tries should it make")
)

func RandomModel() (m1, m2 arith.Model, desc string) {
	switch rand.Intn(1) {
	case 0:
		x1 := arith.Shift2{
			P0: arith.P(rand.Intn(arith.MaxP / 2)),
			I0: byte(rand.Intn(8) + 1),
			P1: arith.P(rand.Intn(arith.MaxP / 2)),
			I1: byte(rand.Intn(8) + 1),
		}
		x2 := x1
		m1, m2 = &x1, &x2
	case 1:
		x1 := arith.Shift4{
			P: [4]arith.P{
				arith.P(rand.Intn(arith.MaxP / 4)),
				arith.P(rand.Intn(arith.MaxP / 4)),
				arith.P(rand.Intn(arith.MaxP / 4)),
				arith.P(rand.Intn(arith.MaxP / 4)),
			},
			I: [4]byte{
				byte(rand.Intn(7) + 1),
				byte(rand.Intn(7) + 1),
				byte(rand.Intn(7) + 1),
				byte(rand.Intn(7) + 1),
			},
		}
		x2 := x1
		m1, m2 = &x1, &x2
	}
	desc = fmt.Sprintf("%#v", m1)
	return
}

func main() {
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	var baseline []physics.Cube
	var current []physics.Cube

	file, err := os.Open(flag.Arg(0))
	check(err)

	dec := gob.NewDecoder(file)
	dec.Decode(&current)
	dec.Decode(&baseline)

	minimal := 1 << 10

	for i := 0; i < *tries; i += 1 {
		menc, _, desc := RandomModel()
		enc := arith.NewEncoder()

		items := []int{}
		for i := range current {
			cube, base := &current[i], &baseline[i]
			if *cube == *base {
				menc.Encode(enc, 0)
			} else {
				menc.Encode(enc, 1)
				items = append(items, i)
			}
		}

		for _, i := range items {
			cube := &current[i]
			menc.Encode(enc, uint(cube.Interacting^1))
		}

		for _, i := range items {
			cube, base := &current[i], &baseline[i]
			v := uint(cube.Largest ^ base.Largest)
			menc.Encode(enc, v&1)
			menc.Encode(enc, v>>1)
		}

		enc.Close()

		if len(enc.Bytes()) < minimal {
			fmt.Printf("%6d %5d   - %s\n", i, len(enc.Bytes()), desc)
			minimal = len(enc.Bytes())
		}
	}
}
