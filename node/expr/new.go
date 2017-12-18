package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type New struct {
	node.SimpleNode
	class     node.Node
	arguments []node.Node
}

func NewNew(class node.Node, arguments []node.Node) node.Node {
	return New{
		node.SimpleNode{Name: "New", Attributes: make(map[string]string)},
		class,
		arguments,
	}
}

func (n New) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.class != nil {
		fmt.Fprintf(out, "\n%vclass:", indent+"  ")
		n.class.Print(out, indent+"    ")
	}

	if n.arguments != nil {
		fmt.Fprintf(out, "\n%varguments:", indent+"  ")
		for _, nn := range n.arguments {
			nn.Print(out, indent+"    ")
		}
	}
}
