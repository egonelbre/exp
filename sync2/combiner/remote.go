package combiner

import (
	"runtime"
	"sync/atomic"
	"unsafe"

	"github.com/egonelbre/exp/sync2/runtime2"
)

type Remote struct {
	batcher Batcher
	done    int64
	_       [7]uint64
	proc    [maxProcessors]remoteProc
}

type remoteNode struct {
	next     unsafe.Pointer // *remoteNode
	argument Argument
}

type remoteProc struct {
	queue int64
	done  int64
	_     [6]uint64
	req   unsafe.Pointer // *remoteNode
	_     [7]uint64
}

func NewRemote(bat Batcher) *Remote {
	return &Remote{batcher: bat}
}

func (c *Remote) Run() {
	var toRelease [maxProcessors]*unsafe.Pointer
	for atomic.LoadInt64(&c.done) == 0 {
		k := 0
		c.batcher.Start()
		for i := 0; i < maxProcessors; i++ {
			proc := &c.proc[i]
			pnode := atomic.LoadPointer(&proc.req)
			if pnode != nil {
				node := (*remoteNode)(pnode)
				c.batcher.Include(node.argument)
				toRelease[k] = &proc.req
				k++
			}
		}
		c.batcher.Finish()
		for _, p := range toRelease[:k] {
			atomic.StorePointer(p, nil)
		}
	}
}

func (c *Remote) Close() { atomic.StoreInt64(&c.done, 1) }

func (c *Remote) Do(arg Argument) {
	node := &remoteNode{argument: arg}

	pid := runtime2.ProcessorHint()
	proc := &c.proc[pid%maxProcessors]

	ticket := atomic.AddInt64(&proc.queue, 1)
	for ticket != atomic.LoadInt64(&proc.done)+1 {
		runtime.Gosched()
	}

	atomic.StorePointer(&proc.req, unsafe.Pointer(node))

	for atomic.LoadPointer(&proc.req) != nil {
		runtime.Gosched()
	}

	atomic.StoreInt64(&proc.done, ticket)
}
