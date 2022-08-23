package main

import (
	"fmt"
	"net/http"
	"reflect"
)

func main() {
	v := reflect.ValueOf(fmt.Println)
	_ = v.MethodByName("Parsed").Call(nil)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
