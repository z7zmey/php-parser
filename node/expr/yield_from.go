package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type YieldFrom struct {
	node.SimpleNode
	expr node.Node
}

func NewYieldFrom(expression node.Node) node.Node {
	return YieldFrom{
		node.SimpleNode{Name: "YieldFrom", Attributes: make(map[string]string)},
		expression,
	}
}

func (n YieldFrom) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
