// implements locking queues, using list and mutex
package goqueuestest

import "sync"

type pcLifoElement struct {
	value interface{}
	prev *pcLifoElement
}

type PcLifo struct {
	head *pcLifoElement
	mutex sync.Mutex
}

func NewPcLifo() *PcLifo {
	q := &PcLifo{}
	q.head = nil
	return q
}

func (q *PcLifo) Enqueue(value interface{}) {
	q.mutex.Lock()
	q.head = &pcLifoElement{value, q.head}
	q.mutex.Unlock()
}

func (q *PcLifo) Dequeue() (interface{}, bool) {
	q.mutex.Lock()
	if q.head == nil {
		q.mutex.Unlock()
		return nil, false
	}
	value := q.head.value
	q.head = q.head.prev
	q.mutex.Unlock()
	return value, true
}

type pcFifoElement struct {
	value interface{}
	next *pcFifoElement
}

type PcFifo struct {
	front *pcFifoElement
	back *pcFifoElement
	mutex sync.Mutex
}

func NewPcFifo() *PcFifo {
	q := &PcFifo{}
	q.front = nil
	q.back = nil
	return q
}

func (q *PcFifo) Enqueue(value interface{}) {
	q.mutex.Lock()
	tmp := &pcFifoElement{value, nil}
	if q.front == nil {
		q.front = tmp
		q.back = tmp
	} else {
		q.back.next = tmp
		q.back = tmp
	}
	q.mutex.Unlock()
}

func (q *PcFifo) Dequeue() (value interface{}, ok bool) {
	q.mutex.Lock()
	tmp := q.front
	if q.front != nil {
		q.front = q.front.next
	} else {
		q.back = nil
	}
	q.mutex.Unlock()
	if tmp != nil {
		return tmp.value, true
	}
	return nil, false
}
