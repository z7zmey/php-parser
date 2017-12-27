package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func(n Require) Name() string {
	return "Require"
}

type Require struct {
	name string
	expr node.Node
}

func NewRequire(expression node.Node) node.Node {
	return Require{
		"Require",
		expression,
	}
}

func (n Require) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
