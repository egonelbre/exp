package reflect2

import (
	"reflect"
	"unsafe"
)

//go:linkname typesByString reflect.typesByString
func typesByString(s string) []unsafe.Pointer

func TypesByString(s string) []reflect.Type {
	var types []reflect.Type
	for _, t := range typesByString(s) {
		pRtypeType := reflect.ValueOf(reflect.TypeOf(0)).Type()
		rtype := reflect.New(pRtypeType).Elem()

		ptr := unsafe.Pointer(reflect.ValueOf(rtype).FieldByName("ptr").Pointer())
		*(*unsafe.Pointer)(ptr) = t

		typ := rtype.Interface().(reflect.Type)
		types = append(types, typ)
	}
	return types
}
