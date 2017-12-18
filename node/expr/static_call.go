package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type StaticCall struct {
	node.SimpleNode
	class     node.Node
	name      node.Node
	arguments []node.Node
}

func NewStaticCall(class node.Node, name node.Node, arguments []node.Node) node.Node {
	return StaticCall{
		node.SimpleNode{Name: "StaticCall", Attributes: make(map[string]string)},
		class,
		name,
		arguments,
	}
}

func (n StaticCall) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.class != nil {
		fmt.Fprintf(out, "\n%vclass:", indent+"  ")
		n.class.Print(out, indent+"    ")
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
