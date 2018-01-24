// implements locking queues, using list and channel
package goqueuestest

import (
	"container/list"
)

type LCFifo struct {
	l    *list.List
	lock chan int
}

func NewListCFifo() *LCFifo {
	ch := make(chan int, 1)
	ch <- 1
	return &LCFifo{list.New(), ch}
}

func (q *LCFifo) Enqueue(value interface{}) {
	<-q.lock
	q.l.PushBack(value)
	q.lock <- 1
}

func (q *LCFifo) Dequeue() (interface{}, bool) {
	<-q.lock
	if q.l.Len() == 0 {
		q.lock <- 1
		return nil, false
	}
	value := q.l.Remove(q.l.Front())
	q.lock <- 1
	return value, true
}

type LCLifo struct {
	l    *list.List
	lock chan int
}

func NewListCLifo() *LCLifo {
	ch := make(chan int, 1)
	ch <- 1
	return &LCLifo{list.New(), ch}
}

func (q *LCLifo) Enqueue(value interface{}) {
	<-q.lock
	q.l.PushBack(value)
	q.lock <- 1
}

func (q *LCLifo) Dequeue() (interface{}, bool) {
	<-q.lock
	if q.l.Len() == 0 {
		q.lock <- 1
		return nil, false
	}
	value := q.l.Remove(q.l.Back())
	q.lock <- 1
	return value, true
}
