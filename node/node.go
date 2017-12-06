package node

import (
	"bytes"
	"fmt"
	"io"
)

type Node interface {
	Print(out io.Writer, indent string)
	Append(nn ...Node) Node
	Attribute(key string, value string) Node
}

type SimpleNode struct {
	Name       string
	Children   []Node
	Attributes map[string]string
}

func (n SimpleNode) String() string {
	buf := new(bytes.Buffer)
	n.Print(buf, " ")
	return buf.String()
}

func (n SimpleNode) Print(out io.Writer, indent string) {
	if len(n.Attributes) > 0 {
		fmt.Fprintf(out, "\n%v%v %s", indent, n.Name, n.Attributes)
	} else {
		fmt.Fprintf(out, "\n%v%v", indent, n.Name)
	}
	for _, nn := range n.Children {
		nn.Print(out, indent+"  ")
	}
}

func NewSimpleNode(name string) Node {
	return SimpleNode{Name: name, Attributes: make(map[string]string)}
}

func (n SimpleNode) Append(nn ...Node) Node {
	n.Children = append(n.Children, nn...)
	return n
}

func (n SimpleNode) Attribute(key string, value string) Node {
	n.Attributes[key] = value
	return n
}
