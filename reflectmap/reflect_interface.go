package reflectmap

import (
	"reflect"
)

// only supports int for key
type Map struct {
	keyType    reflect.Type
	valueType  reflect.Type
	bucketType reflect.Type

	bucketCount int
	buckets     reflect.Value
}

func NewMap(key, value interface{}) *Map {
	keyType := reflect.TypeOf(key)
	valueType := reflect.TypeOf(value)
	bucketCount := 128

	bucketType := reflect.StructOf([]reflect.StructField{
		reflect.StructField{
			Name: "Hash",
			Type: reflect.TypeOf(hashBucket{}),
		},
		reflect.StructField{
			Name: "Key",
			Type: reflect.ArrayOf(bucketSize, keyType),
		},
		reflect.StructField{
			Name: "Value",
			Type: reflect.ArrayOf(bucketSize, valueType),
		},
	})

	buckets := reflect.New(reflect.ArrayOf(bucketCount, bucketType))
	return &Map{
		keyType:    keyType,
		valueType:  valueType,
		bucketType: bucketType,

		bucketCount: bucketCount,
		buckets:     buckets,
	}
}

/*
type bucketType struct {
	hash  [bucketSize]byte
	key   [bucketSize]keyType
	value [bucketSize]valueType
}
*/

func (m *Map) Add(key, value interface{}) {
	// hack to avoid figuring out a hash function
	keyi := key.(int)

	bucketi := keyi % m.bucketCount
	tophash := byte(keyi)
	if tophash == 0 {
		tophash = 1
	}

	bucket := m.buckets.Elem().Index(bucketi)
	hashesField := bucket.Field(0)
	keys := bucket.Field(1)
	values := bucket.Field(2)

	hashes := hashesField.Addr().Interface().(*hashBucket)
	for i, v := range *hashes {
		if v == 0 {
			(*hashes)[i] = tophash
			keys.Index(i).Set(reflect.ValueOf(key))
			values.Index(i).Set(reflect.ValueOf(value))
			return
		}
	}

	(*hashes)[0] = tophash
	keys.Index(0).Set(reflect.ValueOf(key))
	values.Index(0).Set(reflect.ValueOf(value))
}
