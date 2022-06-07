package main

import "math"

// Code based on https://llimllib.github.io/pymag-trees/

func main() {
	/*


	   r = 30
	   rh = r*1.5
	   rw = r*1.5
	   stroke(0)

	   def drawt(root, depth):
	       global r
	       oval(root.x * rw, depth * rh, r, r)
	       print root.x
	       for child in root.children:
	           drawt(child, depth+1)

	   def drawconn(root, depth):
	       for child in root.children:
	           line(root.x * rw + (r/2), depth * rh + (r/2),
	                child.x * rw + (r/2), (depth+1) * rh + (r/2))
	           drawconn(child, depth+1)

	   size(1000, 500)
	   translate(2, 2)
	   stroke(0)
	   drawconn(dt, 0)
	   fill(1,1,1)
	   drawt(dt, 0)
	*/

}

type Node struct {
	X               float32
	Y               int
	Tree            *Node
	Children        []*Node
	Parent          *Node
	Thread          *Node
	Mod             float32
	Ancestor        *Node
	Change, Shift   float32
	LeftMostSibling *Node
	LeftBrother     *Node
	Number          int // 1..n
}

func (node *Node) Init(tree, parent *Node, depth, number int) {
	node.X = -1
	node.Y = depth
	node.Tree = tree
	node.Parent = parent
	node.Ancestor = node
	node.Change = 0
	node.Shift = 0
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

func (node *Node) Right() *Node {
	if node.Thread != nil {
		return node.Thread
	}
	if len(node.Children) > 0 {
		return node.Children[len(node.Children)-1]
	}
	return nil
}

func Buccheim(tree *Node) {
	FirstWalk(tree, 1)
	min := SecondWalk(tree, 0, float32(math.NaN()))
	if min < 0 {
		ThirdWalk(tree, -min)
	}
}

func SecondWalk(n *Node, offset, min float32) float32 {
	n.X += offset
	if min != min || n.X < min {
		min = n.X
	}
	for _, c := range n.Children {
		min = SecondWalk(c, offset+n.Mod, min)
	}
	return min
}

func ThirdWalk(n *Node, offset float32) {
	n.X += offset
	for _, c := range n.Children {
		ThirdWalk(c, offset)
	}
}

func FirstWalk(n *Node, distance float32) {
	if len(n.Children) == 0 {
		if n.LeftMostSibling != nil {
			n.X = n.LeftBrother.X + distance
		} else {
			n.X = 0
		}
		return
	}

	defaultAncestor := n.Children[0]
	for _, c := range n.Children {
		FirstWalk(c, distance)
		defaultAncestor = Apportion(c, defaultAncestor, distance)
	}

	ExecuteShifts(n)

	ell := n.Children[0]
	arr := n.Children[len(n.Children)-1]
	midpoint := (ell.X + arr.X) / 2

	w := n.LeftBrother
	if w != nil {
		n.X = w.X + distance
		n.Mod = n.X + midpoint
	} else {
		n.X = midpoint
	}
}

func ExecuteShifts(n *Node) {
	var shift, change float32
	for i := len(n.Children) - 1; i >= 0; i-- {
		c := n.Children[i]
		c.X += shift
		c.Mod += shift
		change += c.Change
		shift += c.Shift + c.Change
	}
}

func Apportion(v, defaultAncestor *Node, distance float32) *Node {
	w := v.LeftBrother
	if w == nil {
		return defaultAncestor
	}

	vir, vor := v, v
	vil := w
	vol := v.LeftMostSibling
	sir, sor := v.Mod, v.Mod
	sil, sol := vil.Mod, vol.Mod

	for vil.Right() != nil && vir.Left() != nil {
		vil = vil.Right()
		vir = vir.Left()
		vol = vol.Left()
		vor = vor.Right()

		vor.Ancestor = v
		shift := (vil.X + sil) - (vir.X + sir) + distance
		if shift > 0 {
			MoveSubtree(Ancestor(vil, v, defaultAncestor), v, shift)
			sir += shift
			sor += shift
		}
		sil += vil.Mod
		sir += vir.Mod
		sol += vol.Mod
		sor += vor.Mod

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
	}

	return defaultAncestor
}

func MoveSubtree(wl, wr *Node, shift float32) {
	subtrees := float32(wr.Number - wl.Number)
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
