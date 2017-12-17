package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type FunctionCall struct {
	node.SimpleNode
	function  node.Node
	arguments []node.Node
}

func NewFunctionCall(function node.Node, arguments []node.Node) node.Node {
	return FunctionCall{
		node.SimpleNode{Name: "FunctionCall", Attributes: make(map[string]string)},
		function,
		arguments,
	}
}

func (n FunctionCall) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

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
