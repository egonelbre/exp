package arith

const (
	probBits = 12
	MaxP     = 1 << probBits
)

// efficient representation of probability
type P uint32

func Prob(p float64) P { return P(float64(p) * float64(MaxP)) }
