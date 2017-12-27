package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n Property) Name() string {
	return "Property"
}

type Property struct {
	name  string
	token token.Token
	expr  node.Node
}

func NewProperty(token token.Token, expr node.Node) node.Node {
	return Property{
		"Property",
		token,
		expr,
	}
}

func (n Property) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
