package trie

type Uncompact struct {
	root unode
}

type unode struct {
	label  byte
	term   bool
	suffix string
	edges  []unode
}

func (t *Uncompact) Insert(word string) {
	t.root.insert(word)
}

func (n *unode) count() int {
	t := 1
	for i := range n.edges {
		t += n.edges[i].count()
	}
	return t
}

func (n *unode) insert(word string) {
	if word == "" {
		n.term = true
		return
	}

	if len(n.edges) == 0 {
		if n.suffix == "" {
			if safesuffix(word) {
				n.suffix = word
			} else {
				n.insertEdge(word)
			}
		} else {
			n.insertEdge(n.suffix)
			n.insertEdge(word)
			n.suffix = ""
		}
	} else {
		n.insertEdge(word)
	}
}

func (n *unode) insertEdge(word string) {
	p, ok := n.pos(word[0])
	if ok {
		n.edges[p].insert(word[1:])
	} else {
		edge := unode{label: word[0]}
		edge.insert(word[1:])

		n.edges = append(n.edges, unode{})
		copy(n.edges[p+1:], n.edges[p:])
		n.edges[p] = edge
	}
}

func (n *unode) pos(label byte) (int, bool) {
	// TODO: use linear scan
	i, j := 0, len(n.edges)
	for i < j {
		h := i + (j-i)/2
		if n.edges[h].label < label {
			i = h + 1
		} else {
			j = h
		}
	}
	if i < 0 || len(n.edges) <= i {
		return i, false
	}
	return i, n.edges[i].label == label
}

func suffixlen(a, b string) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	if maxsuffix < n {
		n = maxsuffix
	}

	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			return i
		}
	}

	return n
}
