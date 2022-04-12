package timing

import (
	"container/heap"
	"time"
)

type Queue[Key comparable, Value any] struct {
	now     func() time.Duration
	queue   Jobs[Key, Value]
	execute func(Key, Value)
}

type Jobs[Key comparable, Value any] struct {
	items  []Job[Key, Value]
	lookup map[Key]int
}

type Job[Key comparable, Value any] struct {
	deadline time.Duration
	key      Key
	value    Value
}

func (q *Queue[Key, Value]) Len() int {
	return len(q.queue.items)
}

var start = time.Now()

func NewQueue[Key comparable, Value any](exec func(Key, Value)) *Queue[Key, Value] {
	var q Queue[Key, Value]
	q.Init(func() time.Duration { return time.Since(start) }, exec)
	return &q
}

func (q *Queue[Key, Value]) Init(now func() time.Duration, exec func(Key, Value)) {
	q.now = now
	q.queue = Jobs[Key, Value]{
		lookup: map[Key]int{},
	}
	q.execute = exec
}

func (q *Queue[Key, Value]) Add(key Key, value Value, delay time.Duration) {
	heap.Push(&q.queue, Job[Key, Value]{
		deadline: q.now() + delay,
		key:      key,
		value:    value,
	})
}

func (q *Queue[Key, Value]) Move(key Key, delay time.Duration) {
	q.queue.UpdateDeadline(key, q.now() + delay)
}

func (q *Queue[Key, Value]) Remove(key Key) {
	q.queue.Delete(key)
}

/* heap implementation */

func (jobs Jobs[Key, Value]) Len() int { return len(jobs.items) }
func (jobs Jobs[Key, Value]) Less(i, j int) bool { return jobs.items[i].deadline < jobs.items[j].deadline }

func (jobs Jobs[Key, Value]) Swap(i, j int) {
	jobs.items[i], jobs.items[j] = jobs.items[j], jobs.items[i]
	jobs.lookup[jobs.items[i].key] = i
	jobs.lookup[jobs.items[j].key] = j
}

func (jobs *Jobs[Key, Value]) Push(x any) {
	n := len(jobs.items)
	item := x.(Job[Key, Value])
	jobs.lookup[item.key] = n
	jobs.items = append(jobs.items, item)
}

func (jobs *Jobs[Key, Value]) Pop() any {
	old := jobs.items
	n := len(old)
	item := old[n-1]
	old[n-1] = Job[Key, Value]{}
	delete(jobs.lookup, item.key)
	jobs.items = old[0 : n-1]
	return item
}

func (jobs *Jobs[Key, Value]) UpdateDeadline(key Key, deadline time.Duration) {
	index := jobs.lookup[key]
	item := &jobs.items[index]
	item.deadline = deadline
	heap.Fix(jobs, index)
}

func (jobs *Jobs[Key, Value]) Delete(key Key) {
	index := jobs.lookup[key]
	heap.Remove(jobs, index)
}