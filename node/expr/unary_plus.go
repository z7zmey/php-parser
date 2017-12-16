package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type UnaryPlus struct {
	node.SimpleNode
	expr node.Node
}

func NewUnaryPlus(expression node.Node) node.Node {
	return UnaryPlus{
		node.SimpleNode{Name: "UnaryPlus", Attributes: make(map[string]string)},
		expression,
	}
}

func (n UnaryPlus) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
