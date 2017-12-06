package node

import (
	"bytes"
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/token"
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

type tokenNode struct {
	*SimpleNode
	token token.Token
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

func TokenNode(name string, t token.Token) Node {
	return tokenNode{
		&SimpleNode{Name: name, Attributes: make(map[string]string)},
		t,
	}
}

func (n tokenNode) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)
	for _, nn := range n.Children {
		nn.Print(out, indent+"  ")
	}
}

func (n SimpleNode) Append(nn ...Node) Node {
	n.Children = append(n.Children, nn...)
	return n
}

func (n SimpleNode) Attribute(key string, value string) Node {
	n.Attributes[key] = value
	return n
}
