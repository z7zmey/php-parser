package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n Variable) Name() string {
	return "Variable"
}

type Variable struct {
	name     string
	variable node.Node
}

func NewVariable(variable node.Node) node.Node {
	return Variable{
		"Variable",
		variable,
	}
}

func (n Variable) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.variable != nil {
		fmt.Fprintf(out, "\n%vvariable:", indent+"  ")
		n.variable.Print(out, indent+"    ")
	}
}
