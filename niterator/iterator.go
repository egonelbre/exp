package niterator

type Iterator interface {
	IsDone() bool
	Next() (int, error)
}
