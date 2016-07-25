package soap

import "encoding/xml"

type Speced interface {
	Spec() Spec
}

type Spec interface {
	Speced
	Encode() *Node
	Decode(*Node)
}

type Node struct {
	XMLName xml.Name
	Content []byte `xml:",innerxml"`
	Nodes   []Node `xml:",any"`
}

func Parse(content []byte) (*Node, error) {
	var node Node
	err := xml.Unmarshal(content, &node)
	return &node, err
}

func (node *Node) Encode() ([]byte, error) {
	return xml.MarshalIndent(node, "", "\t")
}
