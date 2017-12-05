package name

import (
	"fmt"
	"github.com/z7zmey/php-parser/token"
	"github.com/z7zmey/php-parser/node"
	"io"
)

type NamePart struct {
	node.SimpleNode
	token token.Token
}

func NewNamePart(token token.Token) node.Node {
	return NamePart{
		node.SimpleNode{Name: "NamePart", Attributes: make(map[string]string)}, 
		token,
	}
}

func (n NamePart) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)
	for _, nn := range n.Children {
		nn.Print(out, indent+"  ")
	}
}
