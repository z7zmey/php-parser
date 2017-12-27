package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n Constant) Name() string {
	return "Constant"
}

type Constant struct {
	name  string
	token token.Token
	expr  node.Node
}

func NewConstant(token token.Token, expr node.Node) node.Node {
	return Constant{
		"Constant",
		token,
		expr,
	}
}

func (n Constant) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)

	fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
	n.expr.Print(out, indent+"    ")
}
