package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type ArrayDimFetch struct {
	node.SimpleNode
	variable node.Node
	dim      node.Node
}

func NewArrayDimFetch(variable node.Node, dim node.Node) node.Node {
	return ArrayDimFetch{
		node.SimpleNode{Name: "ArrayDimFetch", Attributes: make(map[string]string)},
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
