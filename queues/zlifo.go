package goqueuestest

import (
	"sync/atomic"
	"unsafe"
)

// Simplistic lock-free stack, suffers from ABA

type ZLifo struct {
	head unsafe.Pointer
}

func NewZLifo() *ZLifo {
	return &ZLifo{nil}
}

func (q *ZLifo) Enqueue(value interface{}) {
	node := unsafe.Pointer(&lfNode{value, nil})
	for {
		(*lfNode)(node).next = q.head
		if atomic.CompareAndSwapPointer(&q.head, (*lfNode)(node).next, node) {
			break
		}
	}
}

func (q *ZLifo) Dequeue() (value interface{}, ok bool) {
	current := q.head
	for current != nil {
		if atomic.CompareAndSwapPointer(&q.head, current, (*lfNode)(current).next) {
			value = (*lfNode)(current).value
			return value, true
		}
		current = q.head
	}
	return nil, false
}
