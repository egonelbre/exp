package fbp

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"sync/atomic"
)

const BufferSize = 0

type Registry map[Type]MakeFn
type MakeFn func() Node

type Graph struct {
	Comm     interface{}
	Registry Registry
	Nodes    map[Name]Node
	Chans    map[string]*Chan
}

type Node interface{}

type Runnable interface {
	Run() error
}

type Chan struct {
	Name string
	Val  reflect.Value
	Refs int32
}

func NewChan(name string, v reflect.Value) *Chan { return &Chan{name, v, 0} }

func (c *Chan) Acquire() { atomic.AddInt32(&c.Refs, 1) }
func (c *Chan) Release() {
	if atomic.AddInt32(&c.Refs, -1) == 0 {
		c.Val.Close()
	}
}

func New(comm interface{}) *Graph {
	return &Graph{
		Comm: comm,

		Registry: make(Registry),
		Nodes:    make(map[Name]Node),
		Chans:    make(map[string]*Chan),
	}
}

func (g *Graph) Setup(def string) error {
	wiring, err := ParseWiring(def)
	if err != nil {
		return err
	}
	return g.WireUp(wiring)
}

func (g *Graph) WireUp(w *Wiring) error {
	g.Nodes["$"] = g.Comm

	for name, typ := range w.Decls {
		mk, ok := g.Registry[typ]
		if !ok {
			return fmt.Errorf("cannot create %s type %s does not exist", name, typ)
		}
		g.Nodes[name] = mk()
	}

	for _, wire := range w.Wires {
		from, ok := g.Nodes[wire.From]
		if !ok {
			return fmt.Errorf("source node %s does not exist", wire.From)
		}

		to, ok := g.Nodes[wire.To]
		if !ok {
			return fmt.Errorf("target node %s does not exist", wire.To)
		}

		rfrom := reflect.ValueOf(from)
		for rfrom.Kind() == reflect.Ptr {
			rfrom = rfrom.Elem()
		}
		rto := reflect.ValueOf(to)
		for rto.Kind() == reflect.Ptr {
			rto = rto.Elem()
		}

		rsrc := rfrom.FieldByName(string(wire.Src))
		rdst := rto.FieldByName(string(wire.Dst))

		if !rsrc.IsValid() {
			return fmt.Errorf("source node %s does not have port %s", wire.From, wire.Src)
		}
		if rsrc.Kind() != reflect.Chan {
			return fmt.Errorf("source %s.%s is not a chan", wire.From, wire.Src)
		}

		if !rdst.IsValid() {
			return fmt.Errorf("target node %s does not have port %s", wire.To, wire.Dst)
		}
		if rdst.Kind() != reflect.Chan {
			return fmt.Errorf("target %s.%s is not a chan", wire.To, wire.Dst)
		}

		srcname := string(wire.From) + "." + string(wire.Src)
		dstname := string(wire.To) + "." + string(wire.Dst)

		var src, dst *Chan

		if rsrc.IsNil() {
			v := reflect.MakeChan(rsrc.Type(), BufferSize)
			rsrc.Set(v)
			src = NewChan(srcname, v)
			g.Chans[srcname] = src
		} else {
			src, ok = g.Chans[srcname]
			if !ok {
				panic("uninitialized src " + srcname)
			}
		}

		if rdst.IsNil() {
			v := reflect.MakeChan(rsrc.Type(), BufferSize)
			rdst.Set(v)
			dst = NewChan(dstname, v)
			g.Chans[dstname] = dst
		} else {
			dst, ok = g.Chans[dstname]
			if !ok {
				panic("uninitialized dst " + dstname)
			}
		}

		dst.Acquire()

		// create a copying routine
		go func() {
			defer dst.Release()
			for {
				v, ok := src.Val.Recv()
				if !ok {
					return
				}
				dst.Val.Send(v)
			}
		}()
	}
	return nil
}

// starts all the nodes
func (g *Graph) Start() {
	for _, n := range g.Nodes {
		r, ok := n.(Runnable)
		if ok {
			go func() {
				//TODO: do something smarter with errors
				if err := r.Run(); err != nil {
					panic(err)
				}
			}()
		}
	}
}

func ParseWiring(def string) (*Wiring, error) {
	wiring := &Wiring{Decls: make(map[Name]Type)}

	// really stupid hacky parsing
	rxDecl := regexp.MustCompile(`:\s+([$a-zA-Z]+)\s+([a-zA-Z]+)`)
	rxPipe := regexp.MustCompile(`([\$a-zA-Z]+)\.([a-zA-Z]+)\s*->\s*([\$a-zA-Z]+)\.([a-zA-Z]+)`)

	line := bufio.NewScanner(bytes.NewBufferString(def))
	for line.Scan() {
		stmt := strings.TrimSpace(line.Text())
		if len(stmt) == 0 {
			continue
		}

		if stmt[0] == ':' {
			xs := rxDecl.FindAllStringSubmatch(stmt, -1)
			if len(xs) != 1 {
				return nil, errors.New("invalid line: " + stmt)
			}

			wiring.Decls[Name(xs[0][1])] = Type(xs[0][2])
		} else {
			xs := rxPipe.FindAllStringSubmatch(stmt, -1)
			if len(xs) != 1 {
				return nil, errors.New("invalid line: " + stmt)
			}

			wiring.Wires = append(wiring.Wires, Wire{
				From: Name(xs[0][1]),
				Src:  Port(xs[0][2]),
				To:   Name(xs[0][3]),
				Dst:  Port(xs[0][4]),
			})
		}
	}

	return wiring, nil
}

type Name string
type Type string
type Port string

type Wiring struct {
	Decls map[Name]Type
	Wires []Wire
}

type Wire struct {
	From Name
	Src  Port
	To   Name
	Dst  Port
}
