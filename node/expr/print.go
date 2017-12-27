package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n Print) Name() string {
	return "Print"
}

type Print struct {
	name string
	expr node.Node
}

func NewPrint(expression node.Node) node.Node {
	return Print{
		"Print",
		expression,
	}
}

func (n Print) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
