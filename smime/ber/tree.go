package ber

import (
	"bytes"
	"fmt"
)

type Tree struct {
	*Token
	Children []*Tree
}

func ParseTree(d *Decoder) (tree *Tree, err error) {
	tree = &Tree{}
	tree.Token, err = d.Token()
	if err != nil {
		return
	}
	if tree.Kind == Constructed {
		var sub *Tree
		for {
			sub, err = ParseTree(d)
			if err != nil {
				return
			}
			if sub.Token.Kind == EndConstructed {
				break
			}

			tree.Children = append(tree.Children, sub)
		}
	}
	return
}

func writeTree(out *forkableWriter, tree *Tree) {
	if len(tree.Children) == 0 {
		h := header{tree.Class, tree.Tag, len(tree.Bytes), false, false}
		marshalHeader(out, h)
		out.Write(tree.Bytes)
	} else {
		head, body := out.fork()
		last := body
		for _, sub := range tree.Children {
			var pre *forkableWriter
			pre, last = last.fork()
			writeTree(pre, sub)
		}

		h := header{tree.Class, tree.Tag, body.Len(), true, false}
		marshalHeader(head, h)
	}
}

func EncodeTree(tree *Tree) ([]byte, error) {
	var out bytes.Buffer
	f := newForkableWriter()
	writeTree(f, tree)
	f.writeTo(&out)
	return out.Bytes(), nil
}

func DecodeTree(data []byte) (*Tree, error) {
	d := NewDecoder(bytes.NewReader(data))
	return ParseTree(d)
}

func (t *Tree) mustNotBeEnd() {
	if t.Kind == EndConstructed {
		panic("invalid tree type")
	}
}

func (t *Tree) AsOctetString() (ret []byte, err error) {
	t.mustNotBeEnd()
	if t.Kind == Value {
		return t.Token.AsOctetString()
	}
	// constructed octet string
	if t.Kind == Constructed {
		var buf []byte
		expectClass := t.Class
		expectTag := t.Tag
		for _, chunk := range t.Children {
			if chunk.Tag != expectTag || chunk.Class != expectClass {
				return nil, fmt.Errorf("invalid constructed octet string")
			}
			buf, err = chunk.AsOctetString()
			if err != nil {
				return
			}
			ret = append(ret, buf...)
		}
	}
	return
}
