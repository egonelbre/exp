package combiner

type Nop struct {
	batcher Batcher
}

func NewNop(batcher Batcher) *Nop { return &Nop{batcher} }

func (nop *Nop) Do(v Argument) {
	nop.batcher.Start()
	nop.batcher.Include(v)
	nop.batcher.Finish()
}
