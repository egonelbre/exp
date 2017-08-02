package view

import "reflect"

type Element interface{}

type Prop struct {
	Type  string
	Value reflect.Value
}

type Node struct {
	Name  string
	Type  string
	Value reflect.Value
	Slice bool

	Props    map[string]*Prop
	Children []*Node
}

func parse(name string, v reflect.Value) (*Node, error) {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct && v.Kind() != reflect.Slice {
		panic("ViewModel should be a struct or slice.")
	}

	t := v.Type()

	node := &Node{
		Name:  name,
		Type:  t.Name(),
		Slice: v.Kind() == reflect.Slice,
		Value: v,
	}

	if node.Type == "" {
		node.Type = node.Name
	}

	if v.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			f := v.Index(i)
			child, err := parse("", f)
			if err != nil {
				return nil, err
			}
			node.Children = append(node.Children, child)
		}
		return node, nil
	}

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Kind() == reflect.Struct || f.Kind() == reflect.Slice {
			child, err := parse(t.Field(i).Name, f)
			if err != nil {
				return nil, err
			}
			node.Children = append(node.Children, child)
		} else {
			if node.Props == nil {
				node.Props = make(map[string]*Prop)
			}

			name := t.Field(i).Name
			node.Props[name] = &Prop{
				Type:  f.Type().Name(),
				Value: f,
			}
		}
	}

	return node, nil
}

func (n *Node) prop(name string) (*Prop, bool) {
	if n.Props == nil {
		return nil, false
	}
	prop, ok := n.Props[name]
	return prop, ok
}

func (n *Node) Str(name string) (string, bool) {
	v, ok := n.prop(name)
	if !ok {
		return "", false
	}
	return v.Value.String(), true
}

func (n *Node) Bool(name string) (bool, bool) {
	v, ok := n.prop(name)
	if !ok {
		return false, false
	}
	return v.Value.Bool(), true
}

func Parse(name string, v interface{}) (*Node, error) {
	return parse(name, reflect.ValueOf(v))
}
