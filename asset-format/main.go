// A memory mappable, zero-copy asset format based on:
//   https://www.youtube.com/watch?v=qZM2B4D7hZs
//
// Warning: using this with mmap is not trivial
//          https://groups.google.com/forum/#!topic/golang-nuts/9iTCLArQmYc
package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

type Ref int32

func (r *Ref) Lock() {}

func (r *Ref) Clear() {
	*r = 0
}
func (r *Ref) Set(v unsafe.Pointer) {
	*r = Ref(uintptr(unsafe.Pointer(v)) - uintptr(unsafe.Pointer(r)))
}
func (r *Ref) Get() unsafe.Pointer {
	return unsafe.Pointer(uintptr(unsafe.Pointer(r)) + uintptr(*r))
}

type Node struct {
	Value  uint32
	_Left  Ref
	_Right Ref
}

func (n *Node) Left() *Node  { return (*Node)(n._Left.Get()) }
func (n *Node) Right() *Node { return (*Node)(n._Right.Get()) }

type Stream struct {
	Data [1 << 10]byte
	Head int
}

func (s *Stream) Write(v interface{}) unsafe.Pointer {
	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)
	size := int(rt.Elem().Size())

	base := unsafe.Pointer(rv.Pointer())
	data := (*[1 << 10]byte)(base)[:size]

	p := unsafe.Pointer(&s.Data[s.Head])
	s.Head += copy(s.Data[s.Head:], data)
	return p
}

func serialize() []byte {
	var stream Stream
	var emptyRef Ref
	index := stream.Write(&emptyRef)
	left := stream.Write(&Node{4129422, 0, 0})
	right := stream.Write(&Node{1238471, 0, 0})
	root := stream.Write(&Node{329501, 0, 0})

	(*Node)(root)._Left.Set(left)
	(*Node)(root)._Right.Set(right)
	(*Ref)(index).Set(root)

	return stream.Data[:stream.Head]
}

func main() {
	data := serialize()
	defer runtime.KeepAlive(data)

	index := (*Ref)(unsafe.Pointer(&data[0]))
	root := (*Node)(index.Get())
	left := root.Left()
	right := root.Right()

	fmt.Println(root)
	fmt.Println(left)
	fmt.Println(right)
}
