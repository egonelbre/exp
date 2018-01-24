// implements locking queues, using list and mutex
package goqueuestest

type ScLifo struct {
	l    []interface{}
	last int
	lock chan int
}

func NewScLifo() *ScLifo {
	q := &ScLifo{}
	q.l = make([]interface{}, 0, growBy)
	q.last = -1
	q.lock = make(chan int, 1)
	q.lock <- 1
	return q
}

func (q *ScLifo) Enqueue(value interface{}) {
	<-q.lock
	q.l = append(q.l, value)
	q.last += 1
	q.lock <- 1
}

func (q *ScLifo) Dequeue() (interface{}, bool) {
	<-q.lock
	if q.last < 0 {
		q.lock <- 1
		return nil, false
	}
	value := q.l[q.last]
	q.last -= 1
	q.lock <- 1
	return value, true
}

type ScFifo struct {
	l      []interface{}
	tail   int
	head   int
	length int
	lock   chan int
}

func NewScFifo() *ScFifo {
	q := &ScFifo{}
	q.l = make([]interface{}, growBy)
	q.tail = 0
	q.head = 0
	q.length = 0
	q.lock = make(chan int, 1)
	q.lock <- 1
	return q
}

func (q *ScFifo) Enqueue(value interface{}) {
	<-q.lock
	if q.length >= len(q.l) {
		q.l = append(q.l[q.tail:], q.l[:q.head]...)
		q.l = append(q.l, make([]interface{}, growBy)...)
		q.tail = 0
		q.head = q.length
	}
	q.l[q.head] = value
	q.head = (q.head + 1) % len(q.l)
	q.length += 1
	q.lock <- 1
}

func (q *ScFifo) Dequeue() (interface{}, bool) {
	<-q.lock
	if q.length == 0 {
		q.lock <- 1
		return nil, false
	}
	value := q.l[q.tail]
	q.tail = (q.tail + 1) % len(q.l)
	q.length -= 1
	q.lock <- 1
	return value, true
}
