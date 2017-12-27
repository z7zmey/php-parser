package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n ArrayDimFetch) Name() string {
	return "ArrayDimFetch"
}

type ArrayDimFetch struct {
	name     string
	variable node.Node
	dim      node.Node
}

func NewArrayDimFetch(variable node.Node, dim node.Node) node.Node {
	return ArrayDimFetch{
		"ArrayDimFetch",
		variable,
		dim,
	}
}

func (n ArrayDimFetch) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.variable != nil {
		fmt.Fprintf(out, "\n%vvariable:", indent+"  ")
		n.variable.Print(out, indent+"    ")
	}

	if n.dim != nil {
		fmt.Fprintf(out, "\n%vdim:", indent+"  ")
		n.dim.Print(out, indent+"    ")
	}
}
