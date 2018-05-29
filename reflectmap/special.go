package reflectmap

type SpecialMap struct {
	bucketCount int
	buckets     []specialBucket
}

type specialBucket struct {
	hash  hashBucket
	key   [bucketSize]int
	value [bucketSize]int
}

func NewSpecialMap() *SpecialMap {
	bucketCount := 128
	return &SpecialMap{
		bucketCount: bucketCount,
		buckets:     make([]specialBucket, bucketCount),
	}
}

func (m *SpecialMap) Add(key, value int) {
	keyi := key

	bucketi := keyi % m.bucketCount
	tophash := byte(keyi)
	if tophash == 0 {
		tophash = 1
	}

	bucket := &m.buckets[bucketi]
	hashes := &bucket.hash
	keys := &bucket.key
	values := &bucket.value

	for i, v := range *hashes {
		if v == 0 {
			(*hashes)[i] = tophash
			(*keys)[i] = key
			(*values)[i] = value
			return
		}
	}

	(*hashes)[0] = tophash
	(*keys)[0] = key
	(*values)[0] = value
}
