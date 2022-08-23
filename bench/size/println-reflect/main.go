package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.ValueOf(fmt.Println)
	_ = v.MethodByName("Parsed").Call(nil)
	fmt.Println("hello, world")
}
