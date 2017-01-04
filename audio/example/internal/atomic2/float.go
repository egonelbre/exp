package atomic2

import (
	"math"
	"sync/atomic"
)

type Float64 struct{ bits uint64 }

func (v *Float64) Get() float64 {
	bits := atomic.LoadUint64(&v.bits)
	return math.Float64frombits(bits)
}

func (v *Float64) Set(to float64) {
	bits := math.Float64bits(to)
	atomic.StoreUint64(&v.bits, bits)
}

func AlmostEqual64(a, b float64) bool {
	const Threshold = 0.001
	d := a - b
	return (-Threshold < d) && (d < Threshold)
}

func AlmostEqual32(a, b float32) bool {
	const Threshold = 0.001
	d := a - b
	return (-Threshold < d) && (d < Threshold)
}
