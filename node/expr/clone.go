package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type Clone struct {
	node.SimpleNode
	expr node.Node
}

func NewClone(expression node.Node) node.Node {
	return Clone{
		node.SimpleNode{Name: "Clone", Attributes: make(map[string]string)},
		expression,
	}
}

func (n Clone) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
