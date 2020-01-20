package cache

type Half struct {
	items       []Entry
	lookup      map[Hash]int
	currentSize int
	maxSize     int
	evicted     EvictFunc
}

func NewHalf(maxSize int, evicted EvictFunc) *Half {
	return &Half{
		items:       []Entry{},
		lookup:      map[Hash]int{},
		currentSize: 0,
		maxSize:     maxSize,
		evicted:     evicted,
	}
}

func NewHalfPrealloc(entries, maxSize int, evicted EvictFunc) *Half {
	return &Half{
		items:       make([]Entry, 0, entries),
		lookup:      make(map[Hash]int, entries),
		currentSize: 0,
		maxSize:     maxSize,
		evicted:     evicted,
	}
}

func (cache *Half) Add(hash Hash, size int) (ok bool) {
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

func (cache *Half) bump(index int) {
	// this could be randomized or use a power of two choices
	cache.swap(index, index/2)
}

func (cache *Half) swap(a, b int) {
	if a == b {
		return
	}

	cache.items[a], cache.items[b] = cache.items[b], cache.items[a]
	cache.lookup[cache.items[a].Hash] = a
	cache.lookup[cache.items[b].Hash] = b
}

func (cache *Half) evict(index int) {
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

func (cache *Half) Get(hash Hash) (e Entry, ok bool) {
	index, ok := cache.lookup[hash]
	if !ok {
		return Entry{}, false
	}

	entry := cache.items[index]
	cache.bump(index)
	return entry, true
}

func (cache *Half) Remove(hash Hash) {
	index, ok := cache.lookup[hash]
	if !ok {
		return
	}
	cache.evict(index)
}

func (cache *Half) Len() int         { return len(cache.items) }
func (cache *Half) CurrentSize() int { return cache.currentSize }
func (cache *Half) MaxSize() int     { return cache.maxSize }
