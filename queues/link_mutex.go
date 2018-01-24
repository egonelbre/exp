// implements locking queues, using list and mutex
package goqueuestest

type pmLifoElement struct {
	value interface{}
	prev *pmLifoElement
}

type PmLifo struct {
	head *pmLifoElement
	lock chan int
}

func NewPmLifo() *PmLifo {
	q := &PmLifo{}
	q.head = nil
	q.lock = make(chan int, 1)
	q.lock <- 1
	return q
}

func (q *PmLifo) Enqueue(value interface{}) {
	<-q.lock
	q.head = &pmLifoElement{value, q.head}
	q.lock <- 1
}

func (q *PmLifo) Dequeue() (interface{}, bool) {
	<-q.lock
	if q.head == nil {
		q.lock <- 1
		return nil, false
	}
	value := q.head.value
	q.head = q.head.prev
	q.lock <- 1
	return value, true
}

type pmFifoElement struct {
	value interface{}
	next *pmFifoElement
}

type PmFifo struct {
	front *pmFifoElement
	back *pmFifoElement
	lock   chan int
}

func NewPmFifo() *PmFifo {
	q := &PmFifo{}
	q.lock = make(chan int, 1)
	q.lock <- 1
	q.front = nil
	q.back = nil
	return q
}

func (q *PmFifo) Enqueue(value interface{}) {
	<-q.lock
	tmp := &pmFifoElement{value, nil}
	if q.front == nil {
		q.front = tmp
		q.back = tmp
	} else {
		q.back.next = tmp
		q.back = tmp
	}
	q.lock <- 1
}

func (q *PmFifo) Dequeue() (value interface{}, ok bool) {
	<-q.lock
	tmp := q.front
	if q.front != nil {
		q.front = q.front.next
	} else {
		q.back = nil
	}
	q.lock <- 1
	if tmp != nil {
		return tmp.value, true
	}
	return nil, false
}
