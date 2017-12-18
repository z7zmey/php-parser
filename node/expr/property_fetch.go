package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type PropertyFetch struct {
	node.SimpleNode
	variable node.Node
	name     node.Node
}

func NewPropertyFetch(variable node.Node, name node.Node) node.Node {
	return PropertyFetch{
		node.SimpleNode{Name: "PropertyFetch", Attributes: make(map[string]string)},
		variable,
		name,
	}
}

func (n PropertyFetch) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.variable != nil {
		fmt.Fprintf(out, "\n%vvariable:", indent+"  ")
		n.variable.Print(out, indent+"    ")
	}

	if n.name != nil {
		fmt.Fprintf(out, "\n%vname:", indent+"  ")
		n.name.Print(out, indent+"    ")
	}
}
