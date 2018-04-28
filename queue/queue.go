// +build ignore

package queue

type Value = int64

type Queue interface {
	// Send puts a value to a queue,
	// returns false when the queue has been closed
	Send(v Value) bool
	// Recv takes a value from the queue
	// returns false when the queue has been closed
	Recv(v *Value) bool
}

type Nonblocking interface {
	// TrySend tries to put a value to a queue
	// returns false when the queue if full or closed
	TrySend(v Value) bool
	// TrySend tries to take a value to a queue
	// returns false when the queue if empty or closed
	TryRecv(v *Value) bool
}

type Closer interface {
	Close()
}

type MPSC interface {
	Queue
	MultipleProducers()
}
type SPMC interface {
	Queue
	MultipleConsumers()
}
type MPMC interface {
	MPSC
	SPMC
}

type NonblockingMPSC interface {
	Nonblocking
	MultipleProducers()
}
type NonblockingSPMC interface {
	Nonblocking
	MultipleConsumers()
}
type NonblockingMPMC interface {
	NonblockingMPSC
	NonblockingSPMC
}
