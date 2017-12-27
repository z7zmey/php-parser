package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n BitwiseNot) Name() string {
	return "BitwiseNot"
}

type BitwiseNot struct {
	name string
	expr node.Node
}

func NewBitwiseNot(expression node.Node) node.Node {
	return BitwiseNot{
		"BitwiseNot",
		expression,
	}
}

func (n BitwiseNot) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
