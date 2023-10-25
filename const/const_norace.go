//go:build !race

package main

func Const[T any](v *T) struct{} { return struct{}{} }
