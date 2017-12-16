package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type BooleanNot struct {
	node.SimpleNode
	expr node.Node
}

func NewBooleanNot(expression node.Node) node.Node {
	return BooleanNot{
		node.SimpleNode{Name: "BooleanNot", Attributes: make(map[string]string)},
		expression,
	}
}

func (n BooleanNot) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
