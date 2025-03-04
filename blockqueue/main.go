package main

import (
	"container/heap"
	"fmt"
)

const BlockSize = 2048

type Element[T any] struct {
	Index    int
	Priority int
	Value    T
}

type Block[T any] [BlockSize]Element[T]

type Queue[T any] struct {
	blocks []*Block[T]
	len    int
}

func (q *Queue[T]) Len() int { return q.len }

func (q *Queue[T]) Push(value T, priority int) {
	heap.Push((*queueHeap[T])(q), Element[T]{
		Index:    0,
		Priority: priority,
		Value:    value,
	})
}

func (q *Queue[T]) Pop() (T, int) {
	e := heap.Pop((*queueHeap[T])(q)).(Element[T])
	return e.Value, e.Priority
}

type queueHeap[T any] Queue[T]

func (q *queueHeap[T]) Len() int {
	return q.len
}

func (q *queueHeap[T]) Less(i, k int) bool {
	ib, ix := i/BlockSize, i%BlockSize
	kb, kx := k/BlockSize, k%BlockSize
	return q.blocks[ib][ix].Priority > q.blocks[kb][kx].Priority
}

func (q *queueHeap[T]) Swap(i, k int) {
	ib, ix := i/BlockSize, i%BlockSize
	kb, kx := k/BlockSize, k%BlockSize

	ie := &q.blocks[ib][ix]
	ke := &q.blocks[kb][kx]

	*ie, *ke = *ke, *ie
	ie.Index = i
	ke.Index = k
}

func (q *queueHeap[T]) Push(x any) {
	e := x.(Element[T])

	ib, ix := q.len/BlockSize, q.len%BlockSize
	if ib >= len(q.blocks) {
		q.blocks = append(q.blocks, new(Block[T]))
	}

	e.Index = q.len
	q.blocks[ib][ix] = e
	q.len++
}

func (q *queueHeap[T]) Pop() any {
	q.len--
	ib, ix := q.len/BlockSize, q.len%BlockSize

	item := q.blocks[ib][ix]
	if ix == 0 {
		q.blocks[ib] = nil
		q.blocks = q.blocks[:len(q.blocks)-1]
	}

	return item
}

func main() {
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	var q Queue[string]

	for value, priority := range items {
		q.Push(value, priority)
	}

	for q.Len() > 0 {
		fmt.Println(q.Pop())
	}
}
