package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type StaticPropertyFetch struct {
	node.SimpleNode
	class node.Node
	name  node.Node
}

func NewStaticPropertyFetch(class node.Node, name node.Node) node.Node {
	return StaticPropertyFetch{
		node.SimpleNode{Name: "StaticPropertyFetch", Attributes: make(map[string]string)},
		class,
		name,
	}
}

func (n StaticPropertyFetch) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.class != nil {
		fmt.Fprintf(out, "\n%vclass:", indent+"  ")
		n.class.Print(out, indent+"    ")
	}

	if n.name != nil {
		fmt.Fprintf(out, "\n%vname:", indent+"  ")
		n.name.Print(out, indent+"    ")
	}
}
