// +build ignore

package main

import (
	"fmt"
	"reflect"

	"github.com/egonelbre/exp/sync2/runtime2/reflect2"
)

func main() {
	for _, t := range reflect2.TypesByString("*runtime.g") {
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}

		goid, ok := t.FieldByName("goid")
		if !ok {
			fmt.Println("did not find g.goid")
		}
		fmt.Printf("#define G_ID_OFFSET %v\n", goid.Offset)

		m, ok := t.FieldByName("m")
		if !ok {
			fmt.Println("did not find g.m")
		}
		fmt.Printf("#define G_M_OFFSET %v\n", m.Offset)

		mp, ok := m.Type.Elem().FieldByName("p")
		if !ok {
			fmt.Println("did not find m.p")
		}
		fmt.Printf("#define M_P_OFFSET %v\n", mp.Offset)
	}

	for _, t := range reflect2.TypesByString("*runtime.p") {
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}

		pid, ok := t.FieldByName("id")
		if !ok {
			fmt.Println("did not find p.id")
		}
		fmt.Printf("#define P_ID_OFFSET %v\n", pid.Offset)
	}
}
