// implements locking queues, using list and mutex
package goqueuestest

import (
	"sync"
)

type SmLifo struct {
	l    []interface{}
	last int
	m    sync.Mutex
}

func NewSmLifo() *SmLifo {
	q := &SmLifo{}
	q.l = make([]interface{}, 0, growBy)
	q.last = -1
	return q
}

func (q *SmLifo) Enqueue(value interface{}) {
	q.m.Lock()
	q.l = append(q.l, value)
	q.last += 1
	q.m.Unlock()
}

func (q *SmLifo) Dequeue() (interface{}, bool) {
	q.m.Lock()
	if q.last < 0 {
		q.m.Unlock()
		return nil, false
	}
	value := q.l[q.last]
	q.last -= 1
	q.m.Unlock()
	return value, true
}

type SmFifo struct {
	l      []interface{}
	tail   int
	head   int
	length int
	m      sync.Mutex
}

func NewSmFifo() *SmFifo {
	q := &SmFifo{}
	q.l = make([]interface{}, growBy)
	q.tail = 0
	q.head = 0
	q.length = 0
	return q
}

func (q *SmFifo) Enqueue(value interface{}) {
	q.m.Lock()
	if q.length >= len(q.l) {
		q.l = append(q.l[q.tail:], q.l[:q.head]...)
		q.l = append(q.l, make([]interface{}, growBy)...)
		q.tail = 0
		q.head = q.length
	}
	q.l[q.head] = value
	q.head = (q.head + 1) % len(q.l)
	q.length += 1
	q.m.Unlock()
}

func (q *SmFifo) Dequeue() (interface{}, bool) {
	q.m.Lock()
	if q.length == 0 {
		q.m.Unlock()
		return nil, false
	}
	value := q.l[q.tail]
	q.tail = (q.tail + 1) % len(q.l)
	q.length -= 1
	q.m.Unlock()
	return value, true
}
