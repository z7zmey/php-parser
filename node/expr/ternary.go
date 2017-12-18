package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type Ternary struct {
	node.SimpleNode
	condition node.Node
	ifTrue    node.Node
	ifFalse   node.Node
}

func NewTernary(condition node.Node, ifTrue node.Node, ifFalse node.Node) node.Node {
	return Ternary{
		node.SimpleNode{Name: "Ternary", Attributes: make(map[string]string)},
		condition,
		ifTrue,
		ifFalse,
	}
}

func (n Ternary) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.condition != nil {
		fmt.Fprintf(out, "\n%vcondition:", indent+"  ")
		n.condition.Print(out, indent+"    ")
	}

	if n.ifTrue != nil {
		fmt.Fprintf(out, "\n%vifTrue:", indent+"  ")
		n.ifTrue.Print(out, indent+"    ")
	}

	if n.ifFalse != nil {
		fmt.Fprintf(out, "\n%vifFalse:", indent+"  ")
		n.ifFalse.Print(out, indent+"    ")
	}
}
