package goqueuestest

import (
	"container/ring"
)

type RcLifo struct {
	head     *ring.Ring // last enqueued value
	length   int
	capacity int
	lock     chan int
}

func NewRcLifo() *RcLifo {
	q := &RcLifo{}
	q.length = 0
	q.capacity = growBy
	q.head = ring.New(growBy)
	q.lock = make(chan int, 1)
	q.lock <- 1
	return q
}

func (q *RcLifo) Enqueue(value interface{}) {
	<-q.lock
	if q.length >= q.capacity {
		q.capacity = q.capacity + growBy
		q.head.Link(ring.New(growBy))
	}
	q.head = q.head.Next()
	q.head.Value = value
	q.length += 1
	q.lock <- 1
}

func (q *RcLifo) Dequeue() (value interface{}, ok bool) {
	<-q.lock
	if q.length == 0 {
		q.lock <- 1
		return nil, false
	}
	value = q.head.Value
	q.head = q.head.Prev()
	q.length -= 1
	q.lock <- 1
	return value, true
}

type RcFifo struct {
	head     *ring.Ring // last enqueued value
	tail     *ring.Ring // last dequeued value
	length   int
	capacity int
	lock     chan int
}

func NewRcFifo() *RcFifo {
	q := &RcFifo{}
	q.length = 0
	q.capacity = growBy
	q.head = ring.New(growBy)
	q.tail = q.head
	q.lock = make(chan int, 1)
	q.lock <- 1
	return q
}

func (q *RcFifo) Enqueue(value interface{}) {
	<-q.lock
	if q.length+1 >= q.capacity {
		q.capacity = q.capacity + growBy
		q.head.Link(ring.New(growBy))
	}
	q.head = q.head.Next()
	q.head.Value = value
	q.length += 1
	q.lock <- 1
}

func (q *RcFifo) Dequeue() (value interface{}, ok bool) {
	<-q.lock
	if q.length == 0 {
		q.lock <- 1
		return nil, false
	}
	q.tail = q.tail.Next()
	value = q.tail.Value
	q.length -= 1
	q.lock <- 1
	return value, true
}
