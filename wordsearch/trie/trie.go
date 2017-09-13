package trie

import "unicode/utf8"

type Iter struct {
	At    *Node
	Index int
}

func (it Iter) Valid() bool { return it.Index >= 0 }
func (it Iter) Match() bool {
	if it.Index == 0 {
		return it.At.Term
	}
	return it.Index == len(it.At.Suffix)
}

func (it Iter) Walk(r rune) Iter {
	if it.At.Suffix != "" {
		sr, sz := utf8.DecodeRuneInString(it.At.Suffix[it.Index:])
		if sr == r {
			return Iter{it.At, it.Index + sz}
		}
		return Iter{nil, -1}
	}

	p, ok := it.At.Pos(r)
	if !ok {
		return Iter{nil, -1}
	}
	return Iter{&it.At.Edges[p], 0}
}

type Node struct {
	Label  rune
	Term   bool
	Suffix string
	Edges  []Node
}

func (n *Node) Insert(word string) {
	if word == "" {
		n.Term = true
		return
	}

	if len(n.Edges) == 0 {
		if n.Suffix == "" {
			n.Suffix = word
		} else {
			n.insertEdge(n.Suffix)
			n.insertEdge(word)
			n.Suffix = ""
		}
	} else {
		n.insertEdge(word)
	}
}

func (n *Node) Contains(word string) bool {
	it := Iter{n, 0}
	if !it.Valid() {
		return false
	}
	for _, r := range word {
		it = it.Walk(r)
		if !it.Valid() {
			return false
		}
	}
	return it.Match()
}

func (n *Node) insertEdge(word string) {
	r, sz := utf8.DecodeRuneInString(word)
	p, ok := n.Pos(r)
	if ok {
		n.Edges[p].Insert(word[sz:])
	} else {
		edge := Node{Label: r}
		edge.Insert(word[sz:])

		n.Edges = append(n.Edges, edge)
		copy(n.Edges[p+1:], n.Edges[p:])
		n.Edges[p] = edge
	}
}

func (n *Node) Pos(label rune) (int, bool) {
	// TODO: use linear scan
	i, j := 0, len(n.Edges)
	for i < j {
		h := i + (j-i)/2
		if n.Edges[h].Label < label {
			i = h + 1
		} else {
			j = h
		}
	}
	if i < 0 || len(n.Edges) <= i {
		return i, false
	}
	return i, n.Edges[i].Label == label
}
