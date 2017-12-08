package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Property struct {
	node.SimpleNode
	token token.Token
	expr  node.Node
}

func NewProperty(token token.Token, expr node.Node) node.Node {
	return Property{
		node.SimpleNode{Name: "Property", Attributes: make(map[string]string)},
		token,
		expr,
	}
}

func (n Property) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
