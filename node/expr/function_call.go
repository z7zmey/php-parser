package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n FunctionCall) Name() string {
	return "FunctionCall"
}

type FunctionCall struct {
	name      string
	function  node.Node
	arguments []node.Node
}

func NewFunctionCall(function node.Node, arguments []node.Node) node.Node {
	return FunctionCall{
		"FunctionCall",
		function,
		arguments,
	}
}

func (n FunctionCall) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.function != nil {
		fmt.Fprintf(out, "\n%vfunction:", indent+"  ")
		n.function.Print(out, indent+"    ")
	}

	if n.arguments != nil {
		fmt.Fprintf(out, "\n%varguments:", indent+"  ")
		for _, nn := range n.arguments {
			nn.Print(out, indent+"    ")
		}
	}
}
