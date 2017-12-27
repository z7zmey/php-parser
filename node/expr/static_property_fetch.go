package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n StaticPropertyFetch) Name() string {
	return "StaticPropertyFetch"
}

type StaticPropertyFetch struct {
	name     string
	class    node.Node
	property node.Node
}

func NewStaticPropertyFetch(class node.Node, property node.Node) node.Node {
	return StaticPropertyFetch{
		"StaticPropertyFetch",
		class,
		property,
	}
}

func (n StaticPropertyFetch) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.class != nil {
		fmt.Fprintf(out, "\n%vclass:", indent+"  ")
		n.class.Print(out, indent+"    ")
	}

	if n.property != nil {
		fmt.Fprintf(out, "\n%vproperty:", indent+"  ")
		n.property.Print(out, indent+"    ")
	}
}
