package goqueuestest

import "unsafe"

type Queue interface {
	Enqueue(value interface{})
	Dequeue() (value interface{}, ok bool)
}

// lockfree node
type lfNode struct {
	value interface{}
	next  unsafe.Pointer
}

const growBy = 1000
