package pathological

import "sync"

const growBy = 1000

type Queue interface {
	Enqueue(value interface{})
	Dequeue() (value interface{}, ok bool)
}

type ScLifo struct {
	l    []interface{}
	last int
	lock chan int
}

func NewScLifo() *ScLifo {
	q := &ScLifo{}
	q.l = make([]interface{}, 0, growBy)
	q.last = -1
	q.lock = make(chan int, 1)
	q.lock <- 1
	return q
}

func (q *ScLifo) Enqueue(value interface{}) {
	<-q.lock
	q.l = append(q.l, value)
	q.last += 1
	q.lock <- 1
}

func (q *ScLifo) Dequeue() (interface{}, bool) {
	<-q.lock
	if q.last < 0 {
		q.lock <- 1
		return nil, false
	}
	value := q.l[q.last]
	q.last -= 1
	q.lock <- 1
	return value, true
}

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