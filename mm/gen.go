// +build ignore

package main

import (
	"os"

	"github.com/egonelbre/exp/mm"
)

func main() {
	data, err := mm.Generate(mm.FallbackDef, map[string]*mm.Spec{
		"Primary": mm.SpecFor(mm.Malloc{}),
		"Fallback": &mm.Spec{
			Type:      "*mm.Region",
			Alignment: mm.PlatformAlignment,
			Dealloc:   true,
			Owns:      true,
			Empty:     false,
		},
	})
	if err != nil {
		panic(err)
	}

	file, err := os.Create("xxx/fallback.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("// DO NOT EDIT\n")
	file.WriteString("// GENERATED CODE\n")
	file.Write(data)
}
