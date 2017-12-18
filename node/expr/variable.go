package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type Variable struct {
	node.SimpleNode
	name node.Node
}

func NewVariable(name node.Node) node.Node {
	return Variable{
		node.SimpleNode{Name: "Variable", Attributes: make(map[string]string)},
		name,
	}
}

func (n Variable) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.name != nil {
		fmt.Fprintf(out, "\n%vname:", indent+"  ")
		n.name.Print(out, indent+"    ")
	}
}
