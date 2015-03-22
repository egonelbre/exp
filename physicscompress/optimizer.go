package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"

	"github.com/egonelbre/exp/coder/arith"
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
	switch rand.Intn(3) {
	case 0:
		x1 := arith.Shift{
			P: arith.Prob(rand.Float64()),
			I: byte(rand.Intn(8)),
		}
		x2 := x1
		m1, m2 = &x1, &x2
	case 1:
		x1 := arith.Shift2{
			P0: arith.Prob(rand.Float64()),
			I0: byte(rand.Intn(10)),
			P1: arith.Prob(rand.Float64()),
			I1: byte(rand.Intn(10)),
		}
		x2 := x1
		m1, m2 = &x1, &x2
	case 2:
		x1 := arith.Shift4{
			P: [4]arith.P{
				arith.P(rand.Intn(arith.MaxP / 4)),
				arith.P(rand.Intn(arith.MaxP / 4)),
				arith.P(rand.Intn(arith.MaxP / 4)),
				arith.P(rand.Intn(arith.MaxP / 4)),
			},
			I: [4]byte{
				byte(rand.Intn(10)),
				byte(rand.Intn(10)),
				byte(rand.Intn(10)),
				byte(rand.Intn(10)),
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

	data, err := ioutil.ReadFile(flag.Arg(0))
	check(err)

	minimal := len(data)
NEXT:
	for i := 0; i < *tries; i += 1 {
		menc, mdec, desc := RandomModel()
		enc := arith.NewEncoder()
		for _, b := range data {
			menc.Encode(enc, uint(b))
		}
		enc.Close()

		dec := arith.NewDecoder(enc.Bytes())
		for _, b := range data {
			v := mdec.Decode(dec)
			if v != uint(b) {
				continue NEXT
			}
		}

		if len(enc.Bytes()) < minimal {
			fmt.Printf("%6d %5d   - %s\n", i, len(enc.Bytes()), desc)
			minimal = len(enc.Bytes())
		}
	}
}
