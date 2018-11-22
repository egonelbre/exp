package iface

import (
	"fmt"
	"testing"
)

func BenchmarkStruct_Direct(b *testing.B) {
	var s = &Struct{10}
	for n := 0; n < b.N; n++ {
		s.x = s.x
		s.x = s.x
		s.x = s.x
		s.x = s.x
		s.x = s.x
		s.x = s.x
		s.x = s.x
		s.x = s.x
		s.x = s.x
		s.x = s.x
	}
}

func BenchmarkStruct_SetDirect(b *testing.B) {
	var s = &Struct{10}
	for n := 0; n < b.N; n++ {
		s.x = 0
		s.x = 1
		s.x = 2
		s.x = 3
		s.x = 4
		s.x = 5
		s.x = 6
		s.x = 7
		s.x = 8
		s.x = 9
	}
}

func BenchmarkStruct_Get(b *testing.B) {
	var s = &Struct{10}
	for n := 0; n < b.N; n++ {
		s.Get()
		s.Get()
		s.Get()
		s.Get()
		s.Get()
		s.Get()
		s.Get()
		s.Get()
		s.Get()
		s.Get()
	}
}

func BenchmarkStruct_Set(b *testing.B) {
	var s = &Struct{10}
	for n := 0; n < b.N; n++ {
		s.Set(n + 0)
		s.Set(n + 1)
		s.Set(n + 2)
		s.Set(n + 3)
		s.Set(n + 4)
		s.Set(n + 5)
		s.Set(n + 6)
		s.Set(n + 7)
		s.Set(n + 8)
		s.Set(n + 9)
	}
}

func BenchmarkStruct_GetCopy(b *testing.B) {
	var s = Struct{10}
	for n := 0; n < b.N; n++ {
		s.CopyGet()
		s.CopyGet()
		s.CopyGet()
		s.CopyGet()
		s.CopyGet()
		s.CopyGet()
		s.CopyGet()
		s.CopyGet()
		s.CopyGet()
		s.CopyGet()
	}
}

func BenchmarkInterface_Get(b *testing.B) {
	s := NewStruct(b.N & 1)
	for n := 0; n < b.N; n++ {
		s.Get()
		s.Get()
		s.Get()
		s.Get()
		s.Get()
		s.Get()
		s.Get()
		s.Get()
		s.Get()
		s.Get()
	}
}

func BenchmarkInterface_CopyGet(b *testing.B) {
	s := NewStruct(b.N & 1)
	for n := 0; n < b.N; n++ {
		s.CopyGet()
		s.CopyGet()
		s.CopyGet()
		s.CopyGet()
		s.CopyGet()
		s.CopyGet()
		s.CopyGet()
		s.CopyGet()
		s.CopyGet()
		s.CopyGet()
	}
}

func BenchmarkInterface_Set(b *testing.B) {
	s := NewStruct(b.N & 1)
	for n := 0; n < b.N; n++ {
		s.Set(0)
		s.Set(0)
		s.Set(0)
		s.Set(0)
		s.Set(0)
		s.Set(0)
		s.Set(0)
		s.Set(0)
		s.Set(0)
		s.Set(0)
	}
}

func BenchmarkSetXReflection(b *testing.B) {
	var s interface{} = &Struct{10}
	//var t *Struct
	for n := 0; n < b.N; n++ {
		switch v := s.(type) {
		case string:
			fmt.Println(v)
		case int32, int64:
			fmt.Println(v)
		case *Struct:
			s.(*Struct).Set(0)
			//t.Set(0)
		default:
			fmt.Printf("unknown %T\n", v)
		}
		switch v := s.(type) {
		case string:
			fmt.Println(v)
		case int32, int64:
			fmt.Println(v)
		case *Struct:
			s.(*Struct).Set(0)
		//t.Set(0)
		default:
			fmt.Printf("unknown %T\n", v)
		}
		switch v := s.(type) {
		case string:
			fmt.Println(v)
		case int32, int64:
			fmt.Println(v)
		case *Struct:
			s.(*Struct).Set(0)
		//t.Set(0)
		default:
			fmt.Printf("unknown %T\n", v)
		}
		switch v := s.(type) {
		case string:
			fmt.Println(v)
		case int32, int64:
			fmt.Println(v)
		case *Struct:
			s.(*Struct).Set(0)
		//t.Set(0)
		default:
			fmt.Printf("unknown %T\n", v)
		}
		switch v := s.(type) {
		case string:
			fmt.Println(v)
		case int32, int64:
			fmt.Println(v)
		case *Struct:
			s.(*Struct).Set(0)
		//t.Set(0)
		default:
			fmt.Printf("unknown %T\n", v)
		}
		switch v := s.(type) {
		case string:
			fmt.Println(v)
		case int32, int64:
			fmt.Println(v)
		case *Struct:
			s.(*Struct).Set(0)
		//t.Set(0)
		default:
			fmt.Printf("unknown %T\n", v)
		}
		switch v := s.(type) {
		case string:
			fmt.Println(v)
		case int32, int64:
			fmt.Println(v)
		case *Struct:
			s.(*Struct).Set(0)
		//t.Set(0)
		default:
			fmt.Printf("unknown %T\n", v)
		}
		switch v := s.(type) {
		case string:
			fmt.Println(v)
		case int32, int64:
			fmt.Println(v)
		case *Struct:
			s.(*Struct).Set(0)
		//t.Set(0)
		default:
			fmt.Printf("unknown %T\n", v)
		}
		switch v := s.(type) {
		case string:
			fmt.Println(v)
		case int32, int64:
			fmt.Println(v)
		case *Struct:
			s.(*Struct).Set(0)
		//t.Set(0)
		default:
			fmt.Printf("unknown %T\n", v)
		}
		switch v := s.(type) {
		case string:
			fmt.Println(v)
		case int32, int64:
			fmt.Println(v)
		case *Struct:
			s.(*Struct).Set(0)
		//t.Set(0)
		default:
			fmt.Printf("unknown %T\n", v)
		}

	}
}
