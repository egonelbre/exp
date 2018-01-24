package goqueuestest

import (
	"container/ring"
	"sync"
)

type RmLifo struct {
	head     *ring.Ring // last enqueued value
	length   int
	capacity int
	m        sync.Mutex
}

func NewRmLifo() *RmLifo {
	q := &RmLifo{}
	q.length = 0
	q.capacity = growBy
	q.head = ring.New(growBy)
	q.m = sync.Mutex{}
	return q
}

func (q *RmLifo) Enqueue(value interface{}) {
	q.m.Lock()
	if q.length >= q.capacity {
		q.capacity = q.capacity + growBy
		q.head.Link(ring.New(growBy))
	}
	q.head = q.head.Next()
	q.head.Value = value
	q.length += 1
	q.m.Unlock()
}

func (q *RmLifo) Dequeue() (value interface{}, ok bool) {
	q.m.Lock()
	if q.length == 0 {
		q.m.Unlock()
		return nil, false
	}
	value = q.head.Value
	q.head = q.head.Prev()
	q.length -= 1
	q.m.Unlock()
	return value, true
}

type RmFifo struct {
	head     *ring.Ring // last enqueued value
	tail     *ring.Ring // last dequeued value
	length   int
	capacity int
	m        sync.Mutex
}

func NewRmFifo() *RmFifo {
	q := &RmFifo{}
	q.length = 0
	q.capacity = growBy
	q.head = ring.New(growBy)
	q.tail = q.head
	q.m = sync.Mutex{}
	return q
}

func (q *RmFifo) Enqueue(value interface{}) {
	q.m.Lock()
	if q.length+1 >= q.capacity {
		q.capacity = q.capacity + growBy
		q.head.Link(ring.New(growBy))
	}
	q.head = q.head.Next()
	q.head.Value = value
	q.length += 1
	q.m.Unlock()
}

func (q *RmFifo) Dequeue() (value interface{}, ok bool) {
	q.m.Lock()
	if q.length == 0 {
		q.m.Unlock()
		return nil, false
	}
	q.tail = q.tail.Next()
	value = q.tail.Value
	q.length -= 1
	q.m.Unlock()
	return value, true
}
