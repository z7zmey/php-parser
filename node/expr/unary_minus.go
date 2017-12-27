package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n UnaryMinus) Name() string {
	return "UnaryMinus"
}

type UnaryMinus struct {
	name string
	expr node.Node
}

func NewUnaryMinus(expression node.Node) node.Node {
	return UnaryMinus{
		"UnaryMinus",
		expression,
	}
}

func (n UnaryMinus) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
