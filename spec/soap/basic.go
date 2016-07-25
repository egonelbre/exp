package soap

type TagSpec struct {
	Name     string
	Children []Speced
}

func Tag(name string, children ...Speced) Spec {
	return &TagSpec{
		Name:     name,
		Children: children,
	}
}

func (spec *TagSpec) Spec() Spec { return spec }

func (spec *TagSpec) Encode() *Node {
	node := &Node{}
	node.XMLName.Local = spec.Name
	for _, child := range spec.Children {
		node.Nodes = append(node.Nodes, *child.Spec().Encode())
	}
	return node
}

func (spec *TagSpec) Decode(node *Node) {
	if spec.Name != node.XMLName.Local {
		panic("invalid name expected " + spec.Name + " got " + node.XMLName.Local)
	}
	if len(spec.Children) != len(node.Nodes) {
		panic("invalid number of children")
	}
	for i, child := range spec.Children {
		child.Spec().Decode(&node.Nodes[i])
	}
}

type StringSpec struct {
	Name  string
	Value *string
}

func String(name string, value *string) Spec {
	return &StringSpec{
		Name:  name,
		Value: value,
	}
}

func (spec *StringSpec) Spec() Spec { return spec }

func (spec *StringSpec) Encode() *Node {
	node := &Node{}
	node.XMLName.Local = spec.Name
	node.Content = []byte(*spec.Value)
	return node
}

func (spec *StringSpec) Decode(node *Node) {
	if spec.Name != node.XMLName.Local {
		panic("invalid name expected " + spec.Name + " got " + node.XMLName.Local)
	}
	if len(node.Nodes) != 0 {
		panic("expected no children for child spec")
	}
	*spec.Value = string(node.Content)
}
