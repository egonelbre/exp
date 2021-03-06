package blockqueue

const (
	BlockSize = 1024
	BlockMask = BlockSize - 1
)

type Queue struct {
	headx  uint32
	tailx  uint32
	head   *Block
	tail   *Block
	blocks uint32
}

type Block struct {
	Items [BlockSize]interface{}
	Next  *Block
}

func New() *Queue { return &Queue{} }

func (q *Queue) Len() int {
	return int(q.blocks)*BlockSize - BlockSize + int(q.tailx) - int(q.headx)
}

func (q *Queue) Empty() bool { return q.head == nil }

func (q *Queue) newBlock() *Block {
	return &Block{}
}

func (q *Queue) tailBlock() *Block {
	if q.tail != nil {
		return q.tail
	}

	q.tail = q.newBlock()
	q.head = q.tail

	return q.tail
}

func (q *Queue) Push(v interface{}) {
	first := q.tailBlock()
	first.Items[q.tailx&BlockMask] = v
	q.tailx++
	if q.tailx == BlockSize {
		q.tailx = 0
		tail := q.tail
		q.tail = q.newBlock()
		tail.Next = q.tail
		q.blocks++
	}
}

func (q *Queue) Pop() (interface{}, bool) {
	if q.head == nil {
		return nil, false
	}

	value := q.head.Items[q.headx&BlockMask]
	q.head.Items[q.headx&BlockMask] = nil
	q.headx++
	if q.headx == BlockSize {
		q.headx = 0
		q.blocks--
		q.head.Next, q.head = nil, q.head.Next
	}

	return value, true
}
