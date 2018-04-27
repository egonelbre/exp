package combiner

type Executor interface {
	Start()
	Include(v int64)
	Finish()
}

type Combiner interface {
	Do(v int64)
}
