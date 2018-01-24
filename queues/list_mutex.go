// implements locking queues, using list and mutex
package goqueuestest

import (
	"container/list"
	"sync"
)

type LFifo struct {
	l *list.List
	m sync.Mutex
}

func NewListFifo() *LFifo {
	return &LFifo{list.New(), sync.Mutex{}}
}

func (q *LFifo) Enqueue(value interface{}) {
	q.m.Lock()
	q.l.PushBack(value)
	q.m.Unlock()
}

func (q *LFifo) Dequeue() (interface{}, bool) {
	q.m.Lock()
	if q.l.Len() == 0 {
		q.m.Unlock()
		return nil, false
	}
	value := q.l.Remove(q.l.Front())
	q.m.Unlock()
	return value, true
}

type LLifo struct {
	l *list.List
	m sync.Mutex
}

func NewListLifo() *LLifo {
	return &LLifo{list.New(), sync.Mutex{}}
}

func (q *LLifo) Enqueue(value interface{}) {
	q.m.Lock()
	q.l.PushBack(value)
	q.m.Unlock()
}

func (q *LLifo) Dequeue() (interface{}, bool) {
	q.m.Lock()
	if q.l.Len() == 0 {
		q.m.Unlock()
		return nil, false
	}
	value := q.l.Remove(q.l.Back())
	q.m.Unlock()
	return value, true
}
