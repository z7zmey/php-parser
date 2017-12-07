package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type Expression struct {
	node.SimpleNode
	expr node.Node
}

func NewExpression(expr node.Node) node.Node {
	return Expression{
		node.SimpleNode{Name: "Expression", Attributes: make(map[string]string)},
		expr,
	}
}

func (n Expression) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
