package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n Throw) Name() string {
	return "Throw"
}

type Throw struct {
	name  string
	token token.Token
	expr  node.Node
}

func NewThrow(token token.Token, expr node.Node) node.Node {
	return Throw{
		"Throw",
		token,
		expr,
	}
}

func (n Throw) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
