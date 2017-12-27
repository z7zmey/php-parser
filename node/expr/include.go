package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n Include) Name() string {
	return "Include"
}

type Include struct {
	name string
	expr node.Node
}

func NewInclude(expression node.Node) node.Node {
	return Include{
		"Include",
		expression,
	}
}

func (n Include) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
