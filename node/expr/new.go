package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n New) Name() string {
	return "New"
}

type New struct {
	name      string
	class     node.Node
	arguments []node.Node
}

func NewNew(class node.Node, arguments []node.Node) node.Node {
	return New{
		"New",
		class,
		arguments,
	}
}

func (n New) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

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
