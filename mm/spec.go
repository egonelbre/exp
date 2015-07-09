package mm

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"text/template"
	"unsafe"
)

type Allocator interface {
	Alignment() int
	Alloc(size int) unsafe.Pointer
}

type Reallocator interface {
	Realloc(p *unsafe.Pointer, size int) bool
}

type Deallocator interface {
	Dealloc(p unsafe.Pointer) bool
}

type Owner interface {
	Owns(p unsafe.Pointer) bool
}

type Emptyer interface {
	Empty() bool
}

type Spec struct {
	Name      string
	Alignment int
	Dealloc   bool
	Owns      bool
	Empty     bool
}

func (s *Spec) String() string { return s.Name }

func SpecFor(v Allocator) *Spec {
	_, dealloc := v.(Deallocator)
	_, owns := v.(Owner)
	_, empty := v.(Emptyer)
	return &Spec{
		Name:      fmt.Sprintf("%+T", v),
		Alignment: v.Alignment(),
		Dealloc:   dealloc,
		Owns:      owns,
		Empty:     empty,
	}
}

var funcs = template.FuncMap{
	"min": func(a, b int) int {
		if a < b {
			return a
		}
		return b
	},
	"max": func(a, b int) int {
		if a > b {
			return a
		}
		return b
	},
}

type Error struct {
	Message   error
	Generated []byte
}

func (err *Error) Error() string {
	if err.Generated == nil {
		return err.Message.Error()
	}
	return fmt.Sprintf("%v\n\n%v", string(err.Generated), err.Message)
}

func Generate(def string, specs map[string]*Spec) ([]byte, error) {
	t := template.Must(template.New("").Funcs(funcs).Parse(def))

	var buf bytes.Buffer
	err := t.Execute(&buf, specs)
	if err != nil {
		return nil, &Error{err, nil}
	}

	data, err := format.Source(buf.Bytes())
	if err != nil {
		os.Stdout.Write(buf.Bytes())
		return nil, &Error{err, nil}
	}

	return data, nil
}
