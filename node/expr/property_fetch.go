package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n PropertyFetch) Name() string {
	return "PropertyFetch"
}

type PropertyFetch struct {
	name     string
	variable node.Node
	property node.Node
}

func NewPropertyFetch(variable node.Node, property node.Node) node.Node {
	return PropertyFetch{
		"PropertyFetch",
		variable,
		property,
	}
}

func (n PropertyFetch) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.variable != nil {
		fmt.Fprintf(out, "\n%vvariable:", indent+"  ")
		n.variable.Print(out, indent+"    ")
	}

	if n.property != nil {
		fmt.Fprintf(out, "\n%vproperty:", indent+"  ")
		n.property.Print(out, indent+"    ")
	}
}
