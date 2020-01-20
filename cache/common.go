package cache

type Hash string

type Entry struct {
	Hash Hash
	Size int
}

type EvictFunc func(Entry)
