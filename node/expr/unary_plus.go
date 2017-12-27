package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n UnaryPlus) Name() string {
	return "UnaryPlus"
}

type UnaryPlus struct {
	name string
	expr node.Node
}

func NewUnaryPlus(expression node.Node) node.Node {
	return UnaryPlus{
		"UnaryPlus",
		expression,
	}
}

func (n UnaryPlus) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
