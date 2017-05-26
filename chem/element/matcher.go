package element

const maxlen = 64

type Matcher struct {
	srclen byte
	src    [maxlen]byte
	trace  [maxlen]Element

	count int
}

func (m *Matcher) Init(word string) {
	copy(m.src[:], []byte(word))
	copy(m.src[len(word):len(word)+5], []byte{0, 0, 0, 0, 0})
	m.srclen = byte(len(word))
	m.count = 0
}

func (m *Matcher) emit(tracelen byte) {
	m.count++
}

func (m *Matcher) Run() { m.acceptElement(0, 0) }

func (m *Matcher) accepted(e Element, p, t byte) {
	m.trace[t] = e
	if p < m.srclen {
		m.acceptElement(p, t+1)
	} else if p == m.srclen {
		m.emit(t)
	}
}
