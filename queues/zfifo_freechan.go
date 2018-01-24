package goqueuestest

import (
	"sync/atomic"
	"unsafe"
)

type ZFifoFreechan struct {
	head     unsafe.Pointer
	tail     unsafe.Pointer
	freelist chan *lfNode
}

// New returns an initialized queue
func NewZFifoFreechan() *ZFifoFreechan {
	q := new(ZFifoFreechan)
	// Creating an initial node
	node := unsafe.Pointer(&lfNode{nil, unsafe.Pointer(q)})

	// Both head and tail point to the initial node
	q.head = node
	q.tail = node

	// freelist for elements
	q.freelist = make(chan *lfNode, 1000)
	return q
}

func (q *ZFifoFreechan) newNode() *lfNode {
	select {
	case element := <-q.freelist:
		return element
	default:
	}
	return &lfNode{}
}

func (q *ZFifoFreechan) freeNode(elem *lfNode) {
	select {
	case q.freelist <- elem:
	default:
	}
}

// Enqueue inserts the value at the tail of the queue
func (q *ZFifoFreechan) Enqueue(value interface{}) {
	node := q.newNode()
	//new(lfNode) // Allocate a new node from the free list
	node.value = value // Copy enqueued value into node
	node.next = unsafe.Pointer(q)
	for { // Keep trying until Enqueue is done
		tail := atomic.LoadPointer(&q.tail)

		// Try to link in new node
		if atomic.CompareAndSwapPointer(&(*lfNode)(tail).next, unsafe.Pointer(q), unsafe.Pointer(node)) {
			// Enqueue is done.  Try to swing tail to the inserted node.
			atomic.CompareAndSwapPointer(&q.tail, tail, unsafe.Pointer(node))
			return
		}

		// Try to swing tail to the next node as the tail was not pointing to the last node
		atomic.CompareAndSwapPointer(&q.tail, tail, (*lfNode)(tail).next)
	}
}

// Dequeue returns the value at the head of the queue and true, or if the queue is empty, it returns a nil value and false
func (q *ZFifoFreechan) Dequeue() (value interface{}, ok bool) {
	for {
		head := atomic.LoadPointer(&q.head)               // Read head pointer
		tail := atomic.LoadPointer(&q.tail)               // Read tail pointer
		next := atomic.LoadPointer(&(*lfNode)(head).next) // Read head.next
		if head != q.head {                               // Check head, tail, and next consistency
			continue // Not consistent. Try again
		}

		if head == tail { // Is queue empty or tail failing behind
			if next == unsafe.Pointer(q) { // Is queue empty?
				return
			}
			// Try to swing tail to the next node as the tail was not pointing to the last node
			atomic.CompareAndSwapPointer(&q.tail, tail, next)
		} else {
			// Read value before CAS
			// Otherwise, another dequeue might free the next node
			value = (*lfNode)(next).value
			// Try to swing Head to the next node
			if atomic.CompareAndSwapPointer(&q.head, head, next) {
				ok = true
				q.freeNode((*lfNode)(head))
				return
			}
			value = nil
		}
	}
	return // Dummy return
}
