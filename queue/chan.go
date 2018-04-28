package queue

type Chan struct {
	ch chan Value
}

func NewChan(size int) *Chan {
	return &Chan{make(chan Value, size)}
}

func (q *Chan) MultipleProducers() {}
func (q *Chan) MultipleConsumers() {}

func (q *Chan) Close() { close(q.ch) }

func (q *Chan) Send(v Value) bool {
	select {
	case q.ch <- v:
		return true
	}
	// return false
}

func (q *Chan) Recv(v *Value) bool {
	select {
	case x := <-q.ch:
		*v = x
		return true
	}
	// return false
}

func (q *Chan) TrySend(v Value) bool {
	select {
	case q.ch <- v:
		return true
	default:
		return false
	}
}

func (q *Chan) TryRecv(v *Value) bool {
	select {
	case x := <-q.ch:
		*v = x
		return true
	default:
		return false
	}
}
