package cache

import (
	"github.com/zeebo/mwc"
)

type Random struct {
	items       []Entry
	lookup      map[Hash]int
	currentSize int
	maxSize     int
	evicted     EvictFunc

	rng mwc.T
}

const (
	lowFactor  = 20
	highFactor = 80
	divisor    = 100
)

func NewRandom(maxSize int, evicted EvictFunc) *Random {
	return &Random{
		items:       []Entry{},
		lookup:      map[Hash]int{},
		currentSize: 0,
		maxSize:     maxSize,
		evicted:     evicted,

		rng: *mwc.Rand(),
	}
}

func NewRandomPrealloc(entries, maxSize int, evicted EvictFunc) *Random {
	return &Random{
		items:       make([]Entry, 0, entries),
		lookup:      make(map[Hash]int, entries),
		currentSize: 0,
		maxSize:     maxSize,
		evicted:     evicted,

		rng: *mwc.Rand(),
	}
}

func (cache *Random) Add(hash Hash, size int) (ok bool) {
	if size > cache.maxSize {
		return false
	}

	cache.currentSize += size
	for cache.currentSize > cache.maxSize {
		cache.evict(len(cache.items) - 1)
	}

	index := len(cache.items)
	cache.lookup[hash] = index
	cache.items = append(cache.items, Entry{hash, size})
	cache.bump(index)

	return true
}

func (cache *Random) bump(index int) {
	low := index * lowFactor / divisor
	high := index * highFactor / divisor
	delta := high - low

	if delta == 0 {
		cache.swap(index, low)
		return
	}

	// we don't need an unbiased random value
	x := cache.rng.Uint64() % uint64(high-low)
	cache.swap(index, low+int(x))
}

func (cache *Random) swap(a, b int) {
	if a == b {
		return
	}

	cache.items[a], cache.items[b] = cache.items[b], cache.items[a]
	cache.lookup[cache.items[a].Hash] = a
	cache.lookup[cache.items[b].Hash] = b
}

func (cache *Random) evict(index int) {
	last := len(cache.items) - 1
	cache.swap(index, last)

	item := cache.items[last]
	cache.items = cache.items[:last]
	cache.currentSize -= item.Size
	delete(cache.lookup, item.Hash)

	if cache.evicted != nil {
		cache.evicted(item)
	}
}

func (cache *Random) Get(hash Hash) (e Entry, ok bool) {
	index, ok := cache.lookup[hash]
	if !ok {
		return Entry{}, false
	}

	entry := cache.items[index]
	cache.bump(index)
	return entry, true
}

func (cache *Random) Remove(hash Hash) {
	index, ok := cache.lookup[hash]
	if !ok {
		return
	}
	cache.evict(index)
}

func (cache *Random) Len() int         { return len(cache.items) }
func (cache *Random) CurrentSize() int { return cache.currentSize }
func (cache *Random) MaxSize() int     { return cache.maxSize }
