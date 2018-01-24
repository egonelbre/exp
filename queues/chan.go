// implements fifo using a channel
package goqueuestest

type CFifo struct {
	ch chan interface{}
}

func NewChanFifo(size int) *CFifo {
	return &CFifo{make(chan interface{}, size)}
}

func (q *CFifo) Enqueue(value interface{}) {
	q.ch <- value
}

func (q *CFifo) Dequeue() (value interface{}, ok bool) {
	select {
	case value, ok = <-q.ch:
		return value, ok
	default:
	}
	return nil, false
}
