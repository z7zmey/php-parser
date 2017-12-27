package assign_op

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type AssignOp struct {
	name       string
	variable   node.Node
	expression node.Node
}

func (n AssignOp) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.variable != nil {
		fmt.Fprintf(out, "\n%vvariable:", indent+"  ")
		n.variable.Print(out, indent+"    ")
	}

	if n.expression != nil {
		fmt.Fprintf(out, "\n%vexpression:", indent+"  ")
		n.expression.Print(out, indent+"    ")
	}
}
