package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"

	"loov.dev/diagram"
)

// Code based on https://llimllib.github.io/pymag-trees/

func RandomTree(depth, width int) *Node {
	root := &Node{}
	if depth <= 0 {
		return root
	}
	n := rand.Intn(width) + 1
	for i := 0; i < n; i++ {
		root.Children = append(root.Children, RandomTree(depth-1, width))
	}
	return root
}

func main() {
	rand.Seed(time.Now().UnixNano())
	root := RandomTree(10, 2)
	root.Init(root, nil, 0, 1)
	// Print(root)
	// fmt.Println()

	Buccheim(root)
	svg := diagram.NewSVG(10000, 10000)
	Draw(svg, root)
	os.WriteFile("tree.svg", svg.Bytes(), 0644)
}

func Print(node *Node) {
	fmt.Printf("Tree(\"%p\"", node)
	for _, child := range node.Children {
		fmt.Print(",\n" + strings.Repeat("\t", node.Y+1))
		Print(child)
	}
	fmt.Printf(")")
}

const (
	R  = 30
	RH = R * 1.5
	RW = R * 1.5
)

func Tree(name string, children ...*Node) *Node {
	return &Node{
		Name:     name,
		Children: children,
	}
}

func Draw(draw diagram.Canvas, n *Node) {
	DrawConn(draw, n)
	DrawNode(draw, n)
}

func Square(cx, cy, d diagram.Length) diagram.Rect {
	return diagram.R(cx-d/2, cy-d/2, cx+d/2, cy+d/2)
}

func DrawNode(draw diagram.Canvas, n *Node) {
	draw.Rect(Square(float64(n.X*RW), float64(n.Y*RH), R), &diagram.Style{
		Stroke: color.Black,
		Size:   1,
		Fill:   color.White,
	})
	for _, c := range n.Children {
		DrawNode(draw, c)
	}
}

func DrawConn(draw diagram.Canvas, n *Node) {
	for _, c := range n.Children {
		draw.Poly(diagram.Ps(
			float64(n.X*RW), float64(n.Y*RH),
			float64(c.X*RW), float64((n.Y+c.Y)*RH)/2,
			float64(c.X*RW), float64(c.Y*RH),
		), &diagram.Style{Stroke: color.Black, Size: 1})
		DrawConn(draw, c)
	}
}

type Node struct {
	Name            string
	X               float32
	Y               int
	Children        []*Node
	Parent          *Node
	Thread          *Node
	Ancestor        *Node
	LeftMostSibling *Node
	LeftBrother     *Node
	Number          int // 1..n

	Mod, Change, Shift float32
}

func (node *Node) Init(tree, parent *Node, depth, number int) {
	node.X = -1
	node.Y = depth
	node.Parent = parent
	node.Ancestor = node
	node.Number = number

	for number, child := range node.Children {
		child.Init(tree, node, depth+1, number+1)
		if number > 0 {
			child.LeftMostSibling = node.Children[0]
			child.LeftBrother = node.Children[number-1]
		}
	}
}

func (node *Node) Left() *Node {
	if node.Thread != nil {
		return node.Thread
	}
	if len(node.Children) > 0 {
		return node.Children[0]
	}
	return nil
}

func (node *Node) String() string {
	return fmt.Sprintf("%v: x=%.1f mod=%v", node.Name, node.X, node.Mod)
}

func (node *Node) Right() *Node {
	if node.Thread != nil {
		return node.Thread
	}
	if len(node.Children) > 0 {
		return node.Children[len(node.Children)-1]
	}
	return nil
}

const Distance = 1

func Buccheim(tree *Node) {
	FirstWalk(tree)
	min := SecondWalk(tree, 0, float32(math.NaN()))
	if min < 0 {
		ThirdWalk(tree, -min)
	}
}

func FirstWalk(n *Node) {
	if len(n.Children) == 0 {
		if n.LeftMostSibling != nil {
			n.X = n.LeftBrother.X + Distance
		} else {
			n.X = 0
		}
		return
	}

	defaultAncestor := n.Children[0]
	for _, c := range n.Children {
		FirstWalk(c)
		defaultAncestor = Apportion(c, defaultAncestor)
	}
	// fmt.Println("finished v =", n, "children")

	ExecuteShifts(n)

	first := n.Children[0]
	last := n.Children[len(n.Children)-1]
	midpoint := (first.X + last.X) / 2

	c := n.LeftBrother
	if c != nil {
		n.X = c.X + Distance
		n.Mod = n.X - midpoint
	} else {
		n.X = midpoint
	}
}

func SecondWalk(n *Node, m, min float32) float32 {
	n.X += m
	if math.IsNaN(float64(min)) || n.X < min {
		min = n.X
	}
	for _, c := range n.Children {
		min = SecondWalk(c, m+n.Mod, min)
	}
	return min
}

func ThirdWalk(n *Node, offset float32) {
	n.X += offset
	for _, c := range n.Children {
		ThirdWalk(c, offset)
	}
}

func ExecuteShifts(n *Node) {
	var shift, change float32
	for i := len(n.Children) - 1; i >= 0; i-- {
		c := n.Children[i]
		// fmt.Println("shift:", c, shift, c.Change)
		c.X += shift
		c.Mod += shift
		change += c.Change
		shift += c.Shift + change
	}
}

func Apportion(v, defaultAncestor *Node) *Node {
	w := v.LeftBrother
	if w == nil {
		return defaultAncestor
	}

	// in buchheim notation:
	// i == inner; o == outer; r == right; l == left; r = +; l = -
	vir, vor := v, v
	vil := w
	vol := v.LeftMostSibling
	sir, sor := v.Mod, v.Mod
	sil := vil.Mod
	sol := vol.Mod

	for vil.Right() != nil && vir.Left() != nil {
		vil = vil.Right()
		vir = vir.Left()
		vol = vol.Left()
		vor = vor.Right()

		vor.Ancestor = v
		shift := (vil.X + sil) - (vir.X + sir) + Distance
		if shift > 0 {
			MoveSubtree(Ancestor(vil, v, defaultAncestor), v, shift)
			sir += shift
			sor += shift
		}
		sil += vil.Mod
		sir += vir.Mod
		sol += vol.Mod
		sor += vor.Mod
	}

	if vil.Right() != nil && vor.Right() == nil {
		vor.Thread = vil.Right()
		vor.Mod += sil - sor
	} else {
		if vir.Left() != nil && vol.Left() == nil {
			vol.Thread = vir.Left()
			vol.Mod += sir - sol
		}
		defaultAncestor = v
	}
	return defaultAncestor
}

func MoveSubtree(wl, wr *Node, shift float32) {
	subtrees := float32(wr.Number - wl.Number)
	// fmt.Println(wl.Name, "is conflicted with", wr.Name, "moving", subtrees, "shift", shift)
	if subtrees <= 0 {
		panic("zero subtrees")
	}
	wr.Change -= shift / subtrees
	wr.Shift += shift
	wl.Change += shift / subtrees
	wr.X += shift
	wr.Mod += shift
}

func Ancestor(vil, v, defaultAncestor *Node) *Node {
	// the relevant text is at the bottom of page 7 of
	// "Improving Walker's Algorithm to Run in Linear Time" by Buchheim et al, (2002)
	// http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.16.8757&rep=rep1&type=pdf
	for _, c := range v.Parent.Children {
		if c == vil.Ancestor {
			return vil.Ancestor
		}
	}
	return defaultAncestor
}
