package queue

import "unsafe"

type Value = int64

const (
	cacheLine = 64
	valueSize = 8
)

// Exposed for intrusive implementations
type Node struct {
	next  unsafe.Pointer
	Value Value
}

type seqValue struct {
	sequence int64
	value    Value
}
type seqPaddedValue struct {
	sequence int64
	value    Value
	_        [8 - 2]int64
}
