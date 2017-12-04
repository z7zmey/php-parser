package scalar

import (
	"fmt"
	"github.com/z7zmey/php-parser/token"
	"github.com/z7zmey/php-parser/node"
	"io"
)

type EncapsedStringPart struct {
	node.SimpleNode
	token token.Token
}


func NewEncapsedStringPart(t token.Token) node.Node {
	return EncapsedStringPart{
		node.SimpleNode{Name: "EncapsedStringPart", Attributes: make(map[string]string)},
		t,
	}
}

func (n EncapsedStringPart) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)
	for _, nn := range n.Children {
		nn.Print(out, indent+"  ")
	}
}
