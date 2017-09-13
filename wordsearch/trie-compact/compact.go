package trie

import (
	"unsafe"
)

const (
	termbit   = 1
	maxsuffix = 6
)

type cindex uint32

type Compact struct {
	nodes []cnode
	root  cnode
}

type cnode struct {
	label   byte
	flags   byte // (edge count << 1) | term
	suffix  [maxsuffix]byte
	edgeptr cindex
}

func (ut *Uncompact) Compress() *Compact {
	t := &Compact{}
	t.nodes = make([]cnode, 0, ut.root.count())
	t.compress(&t.root, &ut.root)
	return t
}

func (t *Compact) compress(dst *cnode, src *unode) {
	dst.label = src.label
	dst.flags = byte(len(src.edges) << 1)
	if src.term {
		dst.flags |= termbit
	}
	copy(dst.suffix[:], []byte(src.suffix))

	dst.edgeptr = cindex(len(t.nodes))
	t.nodes = t.nodes[:len(t.nodes)+len(src.edges) : cap(t.nodes)]
	for i := range src.edges {
		t.compress(t.edge(dst, i), &src.edges[i])
	}
}

func (t *Compact) ContainsBytes(s []byte) bool {
	node := &t.root
next:
	for i, b := range s {
		n := node.edgecount()
		if node.suffix[0] != 0 && len(s)-i <= maxsuffix {
			if node.suffixMatchBytes(s[i:]) {
				return true
			}
		}
		for i := 0; i < n; i++ {
			child := t.edge(node, i)
			if child.label == b {
				node = child
				continue next
			}
		}
		return false
	}
	return node.terminates()
}
func (n *cnode) suffixMatchBytes(s []byte) bool {
	for i, b := range n.suffix {
		if b == 0 {
			return i == len(s)
		}
		if i >= len(s) {
			return false
		}
		if b != s[i] {
			return false
		}
	}
	return true
}

func (t *Compact) Contains(s string) bool {
	node := &t.root
next:
	for i := 0; i < len(s); i++ {
		b := s[i]
		n := node.edgecount()
		if node.suffix[0] != 0 && len(s)-i <= maxsuffix {
			if node.suffixMatch(s[i:]) {
				return true
			}
		}
		for i := 0; i < n; i++ {
			child := t.edge(node, i)
			if child.label == b {
				node = child
				continue next
			}
		}
		return false
	}
	return node.terminates()
}
func (n *cnode) suffixMatch(s string) bool {
	for i, b := range n.suffix {
		if b == 0 {
			return i == len(s)
		}
		if i >= len(s) {
			return false
		}
		if b != s[i] {
			return false
		}
	}
	return true
}
func (t *Compact) edge(n *cnode, i int) *cnode { return &t.nodes[int(n.edgeptr)+i] }

func (n *cnode) edgecount() int   { return int(n.flags >> 1) }
func (n *cnode) terminates() bool { return n.flags&termbit == termbit }

// debugging

func (t *Compact) Size() int { return int(unsafe.Sizeof(cnode{})) * len(t.nodes) }

func (t *Compact) NodeCount() int { return len(t.nodes) }

func (t *Compact) MaxOffset() int {
	max := 0
	for i := range t.nodes {
		dist := int(t.nodes[i].edgeptr) - i
		if dist > max {
			max = dist
		}
	}
	return max
}

func (t *Compact) MaxEdges() int {
	max := 0
	for i := range t.nodes {
		edgelen := int(t.nodes[i].flags >> 1)
		if edgelen > max {
			max = edgelen
		}
	}
	return max
}

func safesuffix(s string) bool {
	if len(s) > maxsuffix {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] == 0 {
			return false
		}
	}
	return true
}
