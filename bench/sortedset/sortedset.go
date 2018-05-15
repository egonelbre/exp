package sortedset

type bucket uint16

const (
	bits = 15        // bits per bucket
	cbit = 1 << bits // continuation bit
	mask = cbit - 1  // bits mask
)

type Set struct {
	cur   int
	count int
	data  []bucket
}

func New() *Set {
	return &Set{}
}

func (s *Set) Add(v int) {
	df := v - s.cur

	if df <= 0 {
		panic("Not in increasing order!")
	}

	s.cur = v
	s.count += 1

	// fast detection of numbits
	switch {
	case df < cbit:
		s.data = append(s.data, bucket(df))

	case df>>(bits*1) < cbit:
		s.data = append(s.data,
			bucket(cbit|(df&mask)),
			bucket(df>>bits),
		)
	case df>>(bits*2) < cbit:
		s.data = append(s.data,
			bucket(cbit|(df&mask)),
			bucket(cbit|((df>>(bits*1))&mask)),
			bucket(df>>(bits*2)),
		)
	case df>>(bits*3) < cbit:
		s.data = append(s.data,
			bucket(cbit|(df&mask)),
			bucket(cbit|((df>>(bits*1))&mask)),
			bucket(cbit|((df>>(bits*2))&mask)),
			bucket(df>>(bits*3)),
		)
	case df>>(bits*4) < cbit:
		s.data = append(s.data,
			bucket(cbit|(df&mask)),
			bucket(cbit|((df>>(bits*1))&mask)),
			bucket(cbit|((df>>(bits*2))&mask)),
			bucket(cbit|((df>>(bits*3))&mask)),
			bucket(df>>(bits*4)),
		)
	default:
		for ; df >= cbit; df = df >> bits {
			s.data = append(s.data, bucket(cbit|(df&mask)))
		}
		s.data = append(s.data, bucket(df))
	}

}

func (s *Set) Len() int {
	return s.count
}

func (s *Set) IterChan(buf int) chan int {
	vals := make(chan int, buf)
	go func() {
		j := 0
		base := 0
		df := 0
		k := uint(0)
		for _, b := range s.data {
			df = df | (int(b&mask) << k)
			k += bits
			if b < cbit {
				base += df
				vals <- base
				df = 0
				k = 0
				j += 1
			}
		}
		close(vals)
	}()
	return vals
}

func (s *Set) IterSlice() []int {
	vals := make([]int, s.count)
	j := 0
	base := 0
	df := 0
	k := uint(0)
	for _, b := range s.data {
		df = df | (int(b&mask) << k)
		k += bits
		if b < cbit {
			base += df
			vals[j] = base
			df = 0
			k = 0
			j += 1
		}
	}
	return vals
}

func (s *Set) IterCallback(fn func(int)) {
	j := 0
	base := 0
	df := 0
	k := uint(0)
	for _, b := range s.data {
		df = df | (int(b&mask) << k)
		k += bits
		if b < cbit {
			base += df
			fn(base)
			df = 0
			k = 0
			j += 1
		}
	}
}

func (s *Set) IterBlockCallback(blockSize int, fn func([]int)) {
	block := make([]int, 0, blockSize)
	j := 0
	base := 0
	df := 0
	k := uint(0)
	for _, b := range s.data {
		df = df | (int(b&mask) << k)
		k += bits
		if b < cbit {
			base += df
			block = append(block, base)
			df = 0
			k = 0
			j += 1
			if cap(block) == len(block) {
				fn(block)
				block = block[0:0:blockSize]
			}
		}
	}
	if len(block) > 0 {
		fn(block)
	}
}
