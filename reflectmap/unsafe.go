package reflectmap

import (
	"reflect"
	"unsafe"
)

// only supports int for key
type UnsafeMap struct {
	keyType    reflect.Type
	valueType  reflect.Type
	bucketType reflect.Type

	keyBase   uintptr
	keySize   uintptr
	valueBase uintptr
	valueSize uintptr

	bucketCount int
	buckets     reflect.Value
	bucketsPtr  uintptr
}

func NewUnsafeMap(key, value interface{}) *UnsafeMap {
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
	return &UnsafeMap{
		keyType:    keyType,
		valueType:  valueType,
		bucketType: bucketType,

		keyBase:   bucketType.FieldByIndex([]int{1}).Offset,
		keySize:   keyType.Size(),
		valueBase: bucketType.FieldByIndex([]int{2}).Offset,
		valueSize: valueType.Size(),

		bucketCount: bucketCount,
		buckets:     buckets,
		bucketsPtr:  buckets.Pointer(),
	}
}

/*
type bucketType struct {
	hash  [bucketSize]byte
	key   [bucketSize]keyType
	value [bucketSize]valueType
}
*/

func (m *UnsafeMap) Add(key, value interface{}) {
	// hack to avoid figuring out a hash function
	keyi := key.(int)

	bucketi := keyi % m.bucketCount
	tophash := byte(keyi)
	if tophash == 0 {
		tophash = 1
	}

	bucketSize := m.bucketType.Size()
	bucketPtr := m.bucketsPtr + bucketSize*uintptr(bucketi)

	hashes := (*hashBucket)(unsafe.Pointer(bucketPtr))
	for i, v := range *hashes {
		if v == 0 {
			(*hashes)[i] = tophash

			keyp := reflect.NewAt(m.keyType, unsafe.Pointer(bucketPtr+m.keyBase+m.keySize*uintptr(i)))
			keyp.Elem().Set(reflect.ValueOf(key))

			valuep := reflect.NewAt(m.valueType, unsafe.Pointer(bucketPtr+m.valueBase+m.valueSize*uintptr(i)))
			valuep.Elem().Set(reflect.ValueOf(value))
			return
		}
	}

	(*hashes)[0] = tophash
	i := 0
	keyp := reflect.NewAt(m.keyType, unsafe.Pointer(bucketPtr+m.keyBase+m.keySize*uintptr(i)))
	keyp.Elem().Set(reflect.ValueOf(key))

	valuep := reflect.NewAt(m.valueType, unsafe.Pointer(bucketPtr+m.valueBase+m.valueSize*uintptr(i)))
	valuep.Elem().Set(reflect.ValueOf(value))
}
