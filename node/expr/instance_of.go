package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type InstanceOf struct {
	node.SimpleNode
	expr  node.Node
	class node.Node
}

func NewInstanceOf(expr node.Node, class node.Node) node.Node {
	return InstanceOf{
		node.SimpleNode{Name: "InstanceOf", Attributes: make(map[string]string)},
		expr,
		class,
	}
}

func (n InstanceOf) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}

	if n.class != nil {
		fmt.Fprintf(out, "\n%vclass:", indent+"  ")
		n.class.Print(out, indent+"    ")
	}
}
