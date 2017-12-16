package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type Require struct {
	node.SimpleNode
	expr node.Node
}

func NewRequire(expression node.Node) node.Node {
	return Require{
		node.SimpleNode{Name: "Require", Attributes: make(map[string]string)},
		expression,
	}
}

func (n Require) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
