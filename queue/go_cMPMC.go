package queue

var _ MPMC = (*MPMCc_go)(nil)

type MPMCc_go struct {
	ch chan Value
}

func NewMPMCc_go(size int) *MPMCc_go {
	return &MPMCc_go{make(chan Value, size)}
}

func (q *MPMCc_go) Cap() int { return cap(q.ch) }

func (q *MPMCc_go) MultipleProducers() {}
func (q *MPMCc_go) MultipleConsumers() {}

func (q *MPMCc_go) Close() { close(q.ch) }

func (q *MPMCc_go) Send(v Value) bool {
	select {
	case q.ch <- v:
		return true
	}
	// return false
}

func (q *MPMCc_go) Recv(v *Value) bool {
	select {
	case x := <-q.ch:
		*v = x
		return true
	}
	// return false
}

func (q *MPMCc_go) TrySend(v Value) bool {
	select {
	case q.ch <- v:
		return true
	default:
		return false
	}
}

func (q *MPMCc_go) TryRecv(v *Value) bool {
	select {
	case x := <-q.ch:
		*v = x
		return true
	default:
		return false
	}
}
