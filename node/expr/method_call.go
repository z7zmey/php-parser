package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type MethodCall struct {
	node.SimpleNode
	variable  node.Node
	name      node.Node
	arguments []node.Node
}

func NewMethodCall(variable node.Node, name node.Node, arguments []node.Node) node.Node {
	return MethodCall{
		node.SimpleNode{Name: "MethodCall", Attributes: make(map[string]string)},
		variable,
		name,
		arguments,
	}
}

func (n MethodCall) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.variable != nil {
		fmt.Fprintf(out, "\n%vvariable:", indent+"  ")
		n.variable.Print(out, indent+"    ")
	}

	if n.name != nil {
		fmt.Fprintf(out, "\n%vname:", indent+"  ")
		n.name.Print(out, indent+"    ")
	}

	if n.arguments != nil {
		fmt.Fprintf(out, "\n%varguments:", indent+"  ")
		for _, nn := range n.arguments {
			nn.Print(out, indent+"    ")
		}
	}
}
