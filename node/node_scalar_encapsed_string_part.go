package node

import (
	"fmt"
	"github.com/z7zmey/php-parser/token"
	"io"
)

type NodeScalarEncapsedStringPart struct {
	*SimpleNode
	token token.Token
}


func NewNodeScalarEncapsedStringPart(t token.Token) Node {
	return NodeScalarEncapsedStringPart{
		&SimpleNode{Name: "NodeScalarEncapsedStringPart", Attributes: make(map[string]string)},
		t,
	}
}

func (n NodeScalarEncapsedStringPart) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)
	for _, nn := range n.Children {
		nn.Print(out, indent+"  ")
	}
}
