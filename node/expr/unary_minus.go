package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type UnaryMinus struct {
	node.SimpleNode
	expr node.Node
}

func NewUnaryMinus(expression node.Node) node.Node {
	return UnaryMinus{
		node.SimpleNode{Name: "UnaryMinus", Attributes: make(map[string]string)},
		expression,
	}
}

func (n UnaryMinus) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
