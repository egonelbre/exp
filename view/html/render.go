package html

import (
	"fmt"
	"io"

	"github.com/egonelbre/exp/view"
)

func Render(root *view.Node, w io.Writer) error {
	fmt.Fprintf("%#v\n", root)
}
