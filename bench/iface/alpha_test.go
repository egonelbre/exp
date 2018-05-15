package iface

import (
	"fmt"
	"testing"
)

func BenchmarkGetXInderectStruct(b *testing.B) {
	var s = &Struct{}
	s.x = 10
	for n := 0; n < b.N; n++ {
		s.GetXIndirect()
		s.GetXIndirect()
		s.GetXIndirect()
		s.GetXIndirect()
		s.GetXIndirect()
		s.GetXIndirect()
		s.GetXIndirect()
		s.GetXIndirect()
		s.GetXIndirect()
		s.GetXIndirect()
	}
}

func BenchmarkGetXManually(b *testing.B) {
	var s = &Struct{}
	s.x = 10
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

func BenchmarkGetXerectStruct(b *testing.B) {
	var s = Struct{}
	s.x = 10
	for n := 0; n < b.N; n++ {
		s.GetX()
		s.GetX()
		s.GetX()
		s.GetX()
		s.GetX()
		s.GetX()
		s.GetX()
		s.GetX()
		s.GetX()
		s.GetX()
	}
}

func BenchmarkGetXInderectInterface(b *testing.B) {
	var s Interface = &Struct{10}

	for n := 0; n < b.N; n++ {
		s.GetXIndirect()
		s.GetXIndirect()
		s.GetXIndirect()
		s.GetXIndirect()
		s.GetXIndirect()
		s.GetXIndirect()
		s.GetXIndirect()
		s.GetXIndirect()
		s.GetXIndirect()
		s.GetXIndirect()

	}
}

func BenchmarkGetXerectInterface(b *testing.B) {
	var s Interface = &Struct{10}
	for n := 0; n < b.N; n++ {
		s.GetX()
		s.GetX()
		s.GetX()
		s.GetX()
		s.GetX()
		s.GetX()
		s.GetX()
		s.GetX()
		s.GetX()
		s.GetX()
	}
}

func BenchmarkSetXInterface(b *testing.B) {
	var s Interface = &Struct{10}
	for n := 0; n < b.N; n++ {
		s.SetX(0)
		s.SetX(0)
		s.SetX(0)
		s.SetX(0)
		s.SetX(0)
		s.SetX(0)
		s.SetX(0)
		s.SetX(0)
		s.SetX(0)
		s.SetX(0)
	}
}

func BenchmarkSetXStruct(b *testing.B) {
	var s = &Struct{10}
	for n := 0; n < b.N; n++ {
		s.SetX(0)
		s.SetX(0)
		s.SetX(0)
		s.SetX(0)
		s.SetX(0)
		s.SetX(0)
		s.SetX(0)
		s.SetX(0)
		s.SetX(0)
		s.SetX(0)
	}
}

func BenchmarkSetXManually(b *testing.B) {
	var s = &Struct{10}
	for n := 0; n < b.N; n++ {
		s.x = 0
		s.x = 0
		s.x = 0
		s.x = 0
		s.x = 0
		s.x = 0
		s.x = 0
		s.x = 0
		s.x = 0
		s.x = 0
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
			s.(*Struct).SetX(0)
			//t.SetX(0)
		default:
			fmt.Printf("unknown %T\n", v)
		}
		switch v := s.(type) {
		case string:
			fmt.Println(v)
		case int32, int64:
			fmt.Println(v)
		case *Struct:
			s.(*Struct).SetX(0)
		//t.SetX(0)
		default:
			fmt.Printf("unknown %T\n", v)
		}
		switch v := s.(type) {
		case string:
			fmt.Println(v)
		case int32, int64:
			fmt.Println(v)
		case *Struct:
			s.(*Struct).SetX(0)
		//t.SetX(0)
		default:
			fmt.Printf("unknown %T\n", v)
		}
		switch v := s.(type) {
		case string:
			fmt.Println(v)
		case int32, int64:
			fmt.Println(v)
		case *Struct:
			s.(*Struct).SetX(0)
		//t.SetX(0)
		default:
			fmt.Printf("unknown %T\n", v)
		}
		switch v := s.(type) {
		case string:
			fmt.Println(v)
		case int32, int64:
			fmt.Println(v)
		case *Struct:
			s.(*Struct).SetX(0)
		//t.SetX(0)
		default:
			fmt.Printf("unknown %T\n", v)
		}
		switch v := s.(type) {
		case string:
			fmt.Println(v)
		case int32, int64:
			fmt.Println(v)
		case *Struct:
			s.(*Struct).SetX(0)
		//t.SetX(0)
		default:
			fmt.Printf("unknown %T\n", v)
		}
		switch v := s.(type) {
		case string:
			fmt.Println(v)
		case int32, int64:
			fmt.Println(v)
		case *Struct:
			s.(*Struct).SetX(0)
		//t.SetX(0)
		default:
			fmt.Printf("unknown %T\n", v)
		}
		switch v := s.(type) {
		case string:
			fmt.Println(v)
		case int32, int64:
			fmt.Println(v)
		case *Struct:
			s.(*Struct).SetX(0)
		//t.SetX(0)
		default:
			fmt.Printf("unknown %T\n", v)
		}
		switch v := s.(type) {
		case string:
			fmt.Println(v)
		case int32, int64:
			fmt.Println(v)
		case *Struct:
			s.(*Struct).SetX(0)
		//t.SetX(0)
		default:
			fmt.Printf("unknown %T\n", v)
		}
		switch v := s.(type) {
		case string:
			fmt.Println(v)
		case int32, int64:
			fmt.Println(v)
		case *Struct:
			s.(*Struct).SetX(0)
		//t.SetX(0)
		default:
			fmt.Printf("unknown %T\n", v)
		}

	}
}
