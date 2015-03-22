package arith

const (
	probBits = 12
	MaxP     = 1 << probBits
)

// efficient representation of probability
type P uint16

func Prob(p float64) P { return P(float64(p) * float64(MaxP)) }

func (p P) midpoint32(lo, hi uint32) uint32 {
	return lo + uint32(uint64(hi-lo)*uint64(p)>>probBits)
}
