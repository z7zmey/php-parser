package assign_op

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type AssignOp struct {
	node.SimpleNode
	variable   node.Node
	expression node.Node
}

func NewAssignOp(variable  node.Node, expression node.Node) node.Node {
	return AssignOp{
		node.SimpleNode{Name: "AssignOp", Attributes: make(map[string]string)},
		variable,
		expression,
	}
}

func (n AssignOp) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.variable != nil {
		fmt.Fprintf(out, "\n%vvariable:", indent+"  ")
		n.variable.Print(out, indent+"    ")
	}
	
	if n.expression != nil {
		fmt.Fprintf(out, "\n%vexpression:", indent+"  ")
		n.expression.Print(out, indent+"    ")
	}
}
