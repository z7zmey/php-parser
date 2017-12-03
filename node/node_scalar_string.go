package node

import (
	"fmt"
	"github.com/z7zmey/php-parser/token"
	"io"
)


type NodeScalarString struct {
	*simpleNode
	token token.Token
}


func NewNodeScalarString(t token.Token) Node {
	return NodeScalarString{
		&simpleNode{name: "NodeScalarString", attributes: make(map[string]string)},
		t,
	}
}

func (n NodeScalarString) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)
	for _, nn := range n.children {
		nn.Print(out, indent+"  ")
	}
}
