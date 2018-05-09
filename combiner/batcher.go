package combiner

type Batcher interface {
	Start()
	Include(op interface{})
	Finish()
}
