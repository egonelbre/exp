package queue

type Value = int64

type Queue interface {
	// BlockingSPSC and/or NonblockingSPSC
	// Closer?

	// Cap should return number of elements that can be sent, before it blocks or fails
	Cap() int
}

type BlockingSPSC interface {
	Queue
	// Send puts a value to a queue,
	// returns false when the queue has been closed
	Send(v Value) bool
	// Recv takes a value from the queue
	// returns false when the queue has been closed
	Recv(v *Value) bool
}

type NonblockingSPSC interface {
	Queue
	// TrySend tries to put a value to a queue
	// returns false when the queue if full or closed
	TrySend(v Value) bool
	// TrySend tries to take a value to a queue
	// returns false when the queue if empty or closed
	TryRecv(v *Value) bool
}

type Flusher interface {
	FlushSend()
	FlushRecv()
}

type Closer interface {
	Close()
}

type BlockingMPSC interface {
	BlockingSPSC
	MultipleProducers()
}

type BlockingSPMC interface {
	BlockingSPSC
	MultipleConsumers()
}

type BlockingMPMC interface {
	BlockingSPSC
	MultipleProducers()
	MultipleConsumers()
}

type NonblockingMPSC interface {
	NonblockingSPSC
	MultipleProducers()
}
type NonblockingSPMC interface {
	NonblockingSPSC
	MultipleConsumers()
}

type NonblockingMPMC interface {
	NonblockingSPSC
	MultipleProducers()
	MultipleConsumers()
}
