package soap

import "reflect"

type TagListSpec struct {
	Name  string
	Items reflect.Value
}

func TagList(name string, items interface{}) Spec {
	v := reflect.ValueOf(items)
	if v.Kind() != reflect.Ptr {
		// todo: better check,
		// i.e. that it's a pointer to elements that implement Spec
		panic("not a pointer to array")
	}
	return &TagListSpec{
		Name:  name,
		Items: v,
	}
}

func (spec *TagListSpec) Spec() Spec { return spec }

//TODO: this can be improved for accepting more different types
func (spec *TagListSpec) Encode() *Node {
	node := &Node{}
	node.XMLName.Local = spec.Name

	items := spec.Items.Elem()
	for i, n := 0, items.Len(); i < n; i++ {
		v := items.Index(i)
		if v.Kind() != reflect.Ptr {
			v = v.Addr()
		}
		child := v.Interface().(Speced)
		node.Nodes = append(node.Nodes, *child.Spec().Encode())
	}

	return node
}

func (spec *TagListSpec) Decode(node *Node) {
	if spec.Name != node.XMLName.Local {
		panic("invalid name expected " + spec.Name + " got " + node.XMLName.Local)
	}
	items := spec.Items.Elem()
	n := len(node.Nodes)

	// Grow slice if necessary
	if n >= items.Cap() {
		newitems := reflect.MakeSlice(items.Type(), items.Len(), n)
		items.Set(newitems)
	}
	items.SetLen(n)

	for i, child := range node.Nodes {
		v := items.Index(i)
		if v.Kind() == reflect.Ptr && v.IsNil() {
			v.Set(reflect.New(v.Type()))
		}
		if v.Kind() != reflect.Ptr {
			v = v.Addr()
		}

		childspec := v.Interface().(Speced).Spec()
		childspec.Decode(&child)
	}
}
