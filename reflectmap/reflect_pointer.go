package reflectmap

import (
	"reflect"
	"unsafe"
)

// only supports int for key
type PointerMap struct {
	keyType    reflect.Type
	valueType  reflect.Type
	bucketType reflect.Type

	tempKey   reflect.Value
	tempValue reflect.Value

	bucketCount int
	buckets     reflect.Value
}

func NewPointerMap(key, value interface{}) *PointerMap {
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
	return &PointerMap{
		keyType:    keyType,
		valueType:  valueType,
		bucketType: bucketType,

		tempKey:   reflect.NewAt(keyType, nil),
		tempValue: reflect.NewAt(valueType, nil),

		bucketCount: bucketCount,
		buckets:     buckets,
	}
}

func (m *PointerMap) Add(key, value unsafe.Pointer) {
	// hack to avoid figuring out a hash function
	keyi := *(*int)(key)

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
			(*reflectvalue)(unsafe.Pointer(&m.tempKey)).ptr = key
			keys.Index(i).Set(m.tempKey.Elem())
			(*reflectvalue)(unsafe.Pointer(&m.tempValue)).ptr = value
			values.Index(i).Set(m.tempValue.Elem())
			return
		}
	}

	i := 0
	(*hashes)[i] = tophash
	(*reflectvalue)(unsafe.Pointer(&m.tempKey)).ptr = key
	keys.Index(i).Set(m.tempKey.Elem())
	(*reflectvalue)(unsafe.Pointer(&m.tempValue)).ptr = value
	values.Index(i).Set(m.tempValue.Elem())
}
