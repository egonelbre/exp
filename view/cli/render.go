package cli

import (
	"fmt"
	"io"
	"strings"

	. "github.com/egonelbre/exp/view"
)

type Lay struct {
	Node     *Node
	Parent   *Lay
	Children []*Lay
}

func Layout(node *Node) (*Lay, error) {
	lay := &Lay{Node: node}
	for _, child := range node.Children {
		sub, err := Layout(child)
		if err != nil {
			return nil, err
		}
		sub.Parent = lay
		lay.Children = append(lay.Children, sub)
	}
	return lay, nil
}

func (lay *Lay) render(w io.Writer, indent string) {
	node := lay.Node

	inslice := lay.Parent != nil && lay.Parent.Node.Slice
	nested := (len(node.Children) > 0)

	name := node.Name
	if name == "" {
		name = node.Type
	}

	switch {
	case inslice:
		fmt.Fprintf(w, "%s┌───────\n", indent)
	case nested:
		fmt.Fprintf(w, "%s┌─ %s ──\n", indent, name)
	}

	switch node.Type {
	case "Field":
		value, _ := node.Str("value")
		back := strings.Repeat("▒", 20)
		fmt.Fprintf(w, "%s%s\n", indent, back+value)
	case "Checkbox":
		value, _ := node.Bool("checked")
		if value {
			fmt.Fprintf(w, "%s│ │\n", indent)
		} else {
			fmt.Fprintf(w, "%s│X│\n", indent)
		}
	case "Button":
		fmt.Fprintf(w, "%s[%s]\n", indent, node.Name)
	}

	for _, child := range lay.Children {
		child.render(w, indent+"  ")
	}

	switch {
	case inslice:
	case nested:
		fmt.Fprintf(w, "%s└──%s───\n", indent, strings.Repeat("─", len(name)))
	}
}

func Render(lay *Lay, w io.Writer) error {
	for _, child := range lay.Children {
		child.render(w, "")
		fmt.Println()
	}
	return nil
}
