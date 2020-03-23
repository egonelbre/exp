package main

import "fmt"

type sortByHash struct {
	keywords []string
	hashes   []byte
}

func (s *sortByHash) Len() int           { return len(s.keywords) }
func (s *sortByHash) Less(i, k int) bool { return s.hashes[i] < s.hashes[k] }
func (s *sortByHash) Swap(i, k int) {
	s.hashes[i], s.hashes[k] = s.hashes[k], s.hashes[i]
	s.keywords[i], s.keywords[k] = s.keywords[k], s.keywords[i]
}

type Combiner struct {
	Name string
	Op   string
	Fn   func(a, b byte) byte
}

var combiners = []*Combiner{
	{Name: "Xor", Op: "^", Fn: func(a, b byte) byte { return a ^ b }},
	{Name: "Add", Op: "+", Fn: func(a, b byte) byte { return a + b }},
	{Name: "Or", Op: "|", Fn: func(a, b byte) byte { return a | b }},
	{Name: "And", Op: "&", Fn: func(a, b byte) byte { return a & b }},
}

func twoHashCalc(name string, a, b *Combiner, shiftn, shift0, shift1 byte) byte {
	return b.Fn(a.Fn(byte(len(name))<<shiftn, name[0]<<shift0), name[1]<<shift1)
}
func twoHashFormat(a, b *Combiner, shiftn, shift0, shift1 byte) string {
	return fmt.Sprintf("((byte(len(name)) << %d) %s (name[0] << %d)) %s (name[1] << %d)", shiftn, a.Op, shift0, b.Op, shift1)
}

func twoHashCalcAlt(name string, a, b *Combiner, shiftn, shift0, shift1 byte) byte {
	return b.Fn(byte(len(name))<<shiftn, a.Fn(name[0]<<shift0, name[1]<<shift1))
}
func twoHashFormatAlt(a, b *Combiner, shiftn, shift0, shift1 byte) string {
	return fmt.Sprintf("(byte(len(name)) << %d) %s ((name[0] << %d) %s (name[1] << %d))", shiftn, a.Op, shift0, b.Op, shift1)
}
