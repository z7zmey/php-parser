package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func(n Expression) Name() string {
	return "Expression"
}

type Expression struct {
	name string
	expr node.Node
}

func NewExpression(expr node.Node) node.Node {
	return Expression{
		"Expression",
		expr,
	}
}

func (n Expression) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
