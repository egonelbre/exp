package imaps

import (
	"strconv"
	"testing"
)

type Alpha struct {
	Name  string
	Value int
	X     int32
	Y     int16
	Z     int64
}

type Beta struct {
	Name  string
	Value int
	X     int32
	Y     int16
	Z     int64
}

type token struct{}

func BenchmarkAlpha(b *testing.B) {
	temp := map[interface{}]token{}
	temp[1] = token{}
	temp[keyAlpha(-1)] = token{}
	for i := 0; i < b.N; i++ {
		_ = temp[keyAlpha(i)]
	}
}

func BenchmarkBeta(b *testing.B) {
	temp := map[interface{}]token{}
	temp[1] = token{}
	temp[keyBeta(-1)] = token{}
	for i := 0; i < b.N; i++ {
		_ = temp[keyBeta(i)]
	}
}
func BenchmarkInt(b *testing.B) {
	temp := map[interface{}]token{}
	temp[1] = token{}
	temp[keyBeta(-1)] = token{}
	for i := 0; i < b.N; i++ {
		_ = temp[strconv.Itoa(i)]
	}
}

func BenchmarkBetaDirect(b *testing.B) {
	temp := map[Beta]token{}
	temp[Beta{}] = token{}
	for i := 0; i < b.N; i++ {
		_ = temp[Beta{Name: "x", Value: i}]
	}
}

//go:noinline
func keyAlpha(v int) interface{} {
	return Alpha{Name: "x", Value: v}
}

//go:noinline
func keyBeta(v int) interface{} {
	return Beta{Name: "x", Value: v}
}
