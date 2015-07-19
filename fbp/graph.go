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
	// Comm is used for input/output from the Graph
	Comm interface{}
	// Registry contains all the constructors for nodes
	Registry Registry
	// Nodes contains created node information
	Nodes map[Name]Node
	// Ports contains all tha ports
	Ports map[string]*Port
}

type Node interface{}

type Runnable interface {
	Run() error
}

type Port struct {
	Name string
	Chan reflect.Value
	Refs int32
}

func NewPort(name string, v reflect.Value) *Port { return &Port{name, v, 0} }

func (c *Port) Acquire() { atomic.AddInt32(&c.Refs, 1) }
func (c *Port) Release() {
	if atomic.AddInt32(&c.Refs, -1) == 0 {
		c.Chan.Close()
	}
}

func New(comm interface{}) *Graph {
	return &Graph{
		Comm: comm,

		Registry: make(Registry),
		Nodes:    make(map[Name]Node),
		Ports:    make(map[string]*Port),
	}
}

// Setup is convenience for parsing the wiring and wiring up the graph
func (g *Graph) Setup(def string) error {
	wiring, err := ParseWiring(def)
	if err != nil {
		return err
	}
	return g.WireUp(wiring)
}

func (g *Graph) WireUp(w *Wiring) error {
	// we add comm as a node to the graph, this simplifies all the wiring
	// I can use $ in the wiring and avoid multiple lookup mechanisms
	g.Nodes["$"] = g.Comm

	// create all the nodes
	for name, typ := range w.Decls {
		mk, ok := g.Registry[typ]
		if !ok {
			return fmt.Errorf("cannot create %s type %s does not exist", name, typ)
		}
		g.Nodes[name] = mk()
	}

	for _, wire := range w.Wires {
		// lookup the source node
		from, ok := g.Nodes[wire.From]
		if !ok {
			return fmt.Errorf("source node %s does not exist", wire.From)
		}

		// lookup the destination node
		to, ok := g.Nodes[wire.To]
		if !ok {
			return fmt.Errorf("target node %s does not exist", wire.To)
		}

		// find the actual source node struct
		rfrom := reflect.ValueOf(from)
		// deref pointers
		// we need to have the actual struct for the FieldByName to work
		for rfrom.Kind() == reflect.Ptr {
			rfrom = rfrom.Elem()
		}
		// same as previous, but for destination node
		rto := reflect.ValueOf(to)
		for rto.Kind() == reflect.Ptr {
			rto = rto.Elem()
		}

		// find the channel fields from nodes
		rsrc := rfrom.FieldByName(string(wire.Src))
		rdst := rto.FieldByName(string(wire.Dst))

		// sanity checks
		switch {
		case !rsrc.IsValid():
			return fmt.Errorf("source node %s does not have port %s", wire.From, wire.Src)
		case rsrc.Kind() != reflect.Chan:
			return fmt.Errorf("source %s.%s is not a chan", wire.From, wire.Src)
		case !rdst.IsValid():
			return fmt.Errorf("target node %s does not have port %s", wire.To, wire.Dst)
		case rdst.Kind() != reflect.Chan:
			return fmt.Errorf("target %s.%s is not a chan", wire.To, wire.Dst)
		}

		// recreate the full port names
		srcname := string(wire.From) + "." + string(wire.Src)
		dstname := string(wire.To) + "." + string(wire.Dst)

		var src, dst *Port
		// have attached the channel to the node already?
		if rsrc.IsNil() {
			// create new channel with correct type
			ch := reflect.MakeChan(rsrc.Type(), BufferSize)
			// add it to the struct
			rsrc.Set(ch)
			// create a port for it
			src = NewPort(srcname, ch)
			// add it to graph
			g.Ports[srcname] = src
		} else {
			// look up the port
			src, ok = g.Ports[srcname]
			// sanity check
			if !ok {
				panic("uninitialized src " + srcname)
			}
		}

		// have attached the channel to the node already?
		if rdst.IsNil() {
			// create new channel with correct type
			ch := reflect.MakeChan(rsrc.Type(), BufferSize)
			// add it to the struct
			rdst.Set(ch)
			// create a port for it
			dst = NewPort(dstname, ch)
			// add it to graph
			g.Ports[dstname] = dst
		} else {
			// look up the port
			dst, ok = g.Ports[dstname]
			// sanity check
			if !ok {
				panic("uninitialized dst " + dstname)
			}
		}

		// we acquire the destination port
		dst.Acquire()

		// create a copying routine
		go func() {
			// when the source finishes we release the destination port
			// this way when the counter hits 0 i.e. there are no more incoming
			// values to the In port of a node then it can be closed
			defer dst.Release()
			for {
				// pull out a value from the output of a node
				v, ok := src.Chan.Recv()
				if !ok {
					return
				}
				// put it into result
				dst.Chan.Send(v)
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
				Src:  PortName(xs[0][2]),
				To:   Name(xs[0][3]),
				Dst:  PortName(xs[0][4]),
			})
		}
	}

	return wiring, nil
}

type Name string
type Type string
type PortName string

type Wiring struct {
	Decls map[Name]Type
	Wires []Wire
}

type Wire struct {
	From Name
	Src  PortName
	To   Name
	Dst  PortName
}
