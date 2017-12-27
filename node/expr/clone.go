package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n Clone) Name() string {
	return "Clone"
}

type Clone struct {
	name string
	expr node.Node
}

func NewClone(expression node.Node) node.Node {
	return Clone{
		"Clone",
		expression,
	}
}

func (n Clone) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
