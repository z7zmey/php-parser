package node

import (
	"bytes"
	"fmt"
	"github.com/z7zmey/php-parser/token"
	"io"
)

type Node interface {
	Print(out io.Writer, indent string)
	Append(nn ...Node) Node
	Attribute(key string, value string) Node
}

type simpleNode struct {
	name       string
	children   []Node
	attributes map[string]string
}

type tokenNode struct {
	*simpleNode
	token token.Token
}

func (n simpleNode) String() string {
	buf := new(bytes.Buffer)
	n.Print(buf, " ")
	return buf.String()
}

func (n simpleNode) Print(out io.Writer, indent string) {
	if len(n.attributes) > 0 {
		fmt.Fprintf(out, "\n%v%v %s", indent, n.name, n.attributes)
	} else {
		fmt.Fprintf(out, "\n%v%v", indent, n.name)
	}
	for _, nn := range n.children {
		nn.Print(out, indent+"  ")
	}
}

func SimpleNode(name string) Node {
	return simpleNode{name: name, attributes: make(map[string]string)}
}

func TokenNode(name string, t token.Token) Node {
	return tokenNode{
		&simpleNode{name: name, attributes: make(map[string]string)},
		t,
	}
}

func (n tokenNode) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)
	for _, nn := range n.children {
		nn.Print(out, indent+"  ")
	}
}

func (n simpleNode) Append(nn ...Node) Node {
	n.children = append(n.children, nn...)
	return n
}

func (n simpleNode) Attribute(key string, value string) Node {
	n.attributes[key] = value
	return n
}
