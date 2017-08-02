package xml

import (
	"fmt"
	"io"
	"strings"

	"github.com/egonelbre/exp/view"
)

type renderer struct{ io.Writer }

func (r renderer) Render(indent string, node *view.Node) error {
	fmt.Fprintf(r, "%s<%s", indent, node.Type)
	if node.Name != "" {
		fmt.Fprintf(r, " class='%s'", node.Name)
	}
	for _, prop := range node.Props {
		fmt.Fprintf(r, " %s='%s'", strings.ToLower(prop.Name), prop.Value)
	}
	fmt.Fprintf(r, ">")
	if len(node.Children) > 0 {
		fmt.Fprintln(r)
	}
	for _, child := range node.Children {
		r.Render(indent+"\t", child)
	}
	if len(node.Children) > 0 {
		fmt.Fprint(r, indent)
	}
	fmt.Fprintf(r, "<\\%s>\n", node.Type)
	return nil
}

func Render(root *view.Node, w io.Writer) error {
	return renderer{w}.Render("", root)
}
