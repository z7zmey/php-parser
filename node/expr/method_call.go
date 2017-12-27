package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n MethodCall) Name() string {
	return "MethodCall"
}

type MethodCall struct {
	name      string
	variable  node.Node
	method    node.Node
	arguments []node.Node
}

func NewMethodCall(variable node.Node, method node.Node, arguments []node.Node) node.Node {
	return MethodCall{
		"MethodCall",
		variable,
		method,
		arguments,
	}
}

func (n MethodCall) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.variable != nil {
		fmt.Fprintf(out, "\n%vvariable:", indent+"  ")
		n.variable.Print(out, indent+"    ")
	}

	if n.method != nil {
		fmt.Fprintf(out, "\n%vmethod:", indent+"  ")
		n.method.Print(out, indent+"    ")
	}

	if n.arguments != nil {
		fmt.Fprintf(out, "\n%varguments:", indent+"  ")
		for _, nn := range n.arguments {
			nn.Print(out, indent+"    ")
		}
	}
}
