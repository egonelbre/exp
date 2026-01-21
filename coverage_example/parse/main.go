package main

import (
	"cmp"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type Module struct {
	Prefix string
	Dir    string
}

// See https://go.dev/ref/mod#go-mod-file for the spec.
func ParseModule(path string) (Module, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Module{}, err
	}

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		// is it a module defintion line?
		if strings.HasPrefix(line, "module ") {
			// trim the prefix
			modulename := strings.TrimPrefix(line, "module ")
			// cut any trailing comments, if there are any
			modulename, _, _ = strings.Cut(modulename, "//")
			modulename = strings.TrimSpace(modulename)

			return Module{
				Prefix: modulename + "/",
				Dir:    filepath.Dir(path),
			}, nil
		}
	}

	return Module{}, errors.New("invalid module")
}

func TranslatePath(pkgpath string, modules []Module) (path string, found bool) {
	for _, module := range modules {
		if strings.HasPrefix(pkgpath, module.Prefix) {
			f := filepath.Join(module.Dir, strings.TrimPrefix(pkgpath, module.Prefix))
			return f, true
		}
	}
	return pkgpath, false
}

func main() {
	modules := []Module{
		must(ParseModule("example/go.mod")),
		must(ParseModule("example/hack/json/go.mod")),
	}

	// sort in reverse prefix length order,
	// the most precise prefix needs to match first
	slices.SortFunc(modules, func(a, b Module) int {
		return -cmp.Compare(len(a.Prefix), len(b.Prefix))
	})

	fmt.Println(TranslatePath("hello/world/main.go", modules))
	fmt.Println(TranslatePath("hello/world/json/main.go", modules))
	fmt.Println(TranslatePath("github.com/external/counter.go", modules))
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
