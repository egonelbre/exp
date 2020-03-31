package lockset

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

// TODO:
//   * per goroutine table
//   * across goroutine locks
//   * replacement for sync package
//   * reduce overhead
//   * garbage collection of locks

type Table struct {
	mu sync.Mutex

	root    Lock
	holding []*Lock
	callers []CallFrame
}

type CallFrame struct {
	entries [4]uintptr
	count   int
}

func (cf *CallFrame) String() string {
	frames := runtime.CallersFrames(cf.entries[:cf.count])
	var builder strings.Builder
	for {
		frame, more := frames.Next()
		fmt.Fprintf(&builder, "    %s (%s:%d)\n", frame.Function, frame.File, frame.Line)
		if !more {
			break
		}
	}
	return builder.String()
}

type Lock struct {
	followers Set
}

func NewTable() *Table {
	table := &Table{}
	table.holding = append(table.holding, &table.root)
	table.callers = append(table.callers, CallFrame{})
	return table
}

type Inversion struct {
	Held     *Lock
	HeldCall CallFrame

	Lock     *Lock
	LockCall CallFrame
}

func (inv *Inversion) String() string {
	return fmt.Sprintf("previously locked:\n%s\nlocking:\n%s", &inv.HeldCall, &inv.LockCall)
}

func (table *Table) Locked(lock *Lock) *Inversion {
	table.mu.Lock()
	defer table.mu.Unlock()

	var frame CallFrame
	frame.count = runtime.Callers(2, frame.entries[:])

	var inversion *Inversion

	for i, held := range table.holding {
		if lock.followers.Contains(held) {
			inversion = &Inversion{
				Held:     held,
				HeldCall: table.callers[i],

				Lock:     lock,
				LockCall: frame,
			}
			break
		}
		held.followers.Include(lock)
	}

	table.holding = append(table.holding, lock)
	table.callers = append(table.callers, frame)

	return inversion
}

func (table *Table) Unlocking(lock *Lock) bool {
	table.mu.Lock()
	defer table.mu.Unlock()

	for i := len(table.holding) - 1; i >= 0; i-- {
		if table.holding[i] == lock {
			table.holding = append(table.holding[:i], table.holding[i+1:]...)
			table.callers = append(table.callers[:i], table.callers[i+1:]...)
			return true
		}
	}

	return false
}

type Set struct{ list []*Lock }

func (set *Set) Include(lock *Lock) {
	if !set.Contains(lock) {
		set.list = append(set.list, lock)
	}
}

func (set *Set) Contains(lock *Lock) bool {
	for _, k := range set.list {
		if k == lock {
			return true
		}
	}
	return false
}

func (set *Set) Remove(lock *Lock) bool {
	for i, k := range set.list {
		if k == lock {
			set.list = append(set.list[:i], set.list[i+1:]...)
			return true
		}
	}
	return false
}
