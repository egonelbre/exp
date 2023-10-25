//go:build race

package main

import (
	"reflect"
	"runtime"
)

func Const[T any](v *T) struct{} {
	go deepRaceRead(reflect.ValueOf(v))
	return struct{}{}
}

func deepRaceRead(v reflect.Value) {
	_ = runtime.KeepAlive
	switch v.Type().Kind() {
	default:
		panic("unhandled type")
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128:
		runtime.RaceRead(v.Addr().UnsafePointer())
	case reflect.Ptr:
		if v.IsNil() {
			return
		}
		deepRaceRead(v.Elem())
	case reflect.String:
		runtime.RaceRead(v.UnsafePointer())
	case reflect.Array, reflect.Slice:
		// TODO: use runtime.RaceReadRange
		for i := 0; i < v.Len(); i++ {
			deepRaceRead(v.Index(i))
		}
	case reflect.Map:
		if v.IsNil() {
			return
		}
		iter := v.MapRange()
		for iter.Next() {
			deepRaceRead(iter.Value())
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			deepRaceRead(v.Field(i))
		}
	}
}
