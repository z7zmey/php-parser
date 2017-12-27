package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n StaticCall) Name() string {
	return "StaticCall"
}

type StaticCall struct {
	name      string
	class     node.Node
	call      node.Node
	arguments []node.Node
}

func NewStaticCall(class node.Node, call node.Node, arguments []node.Node) node.Node {
	return StaticCall{
		"StaticCall",
		class,
		call,
		arguments,
	}
}

func (n StaticCall) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.class != nil {
		fmt.Fprintf(out, "\n%vclass:", indent+"  ")
		n.class.Print(out, indent+"    ")
	}

	if n.call != nil {
		fmt.Fprintf(out, "\n%vcall:", indent+"  ")
		n.call.Print(out, indent+"    ")
	}

	if n.arguments != nil {
		fmt.Fprintf(out, "\n%varguments:", indent+"  ")
		for _, nn := range n.arguments {
			nn.Print(out, indent+"    ")
		}
	}
}
