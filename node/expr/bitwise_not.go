package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type BitwiseNot struct {
	node.SimpleNode
	expr node.Node
}

func NewBitwiseNot(expression node.Node) node.Node {
	return BitwiseNot{
		node.SimpleNode{Name: "BitwiseNot", Attributes: make(map[string]string)},
		expression,
	}
}

func (n BitwiseNot) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
