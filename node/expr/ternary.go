package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n Ternary) Name() string {
	return "Ternary"
}

type Ternary struct {
	name      string
	condition node.Node
	ifTrue    node.Node
	ifFalse   node.Node
}

func NewTernary(condition node.Node, ifTrue node.Node, ifFalse node.Node) node.Node {
	return Ternary{
		"Ternary",
		condition,
		ifTrue,
		ifFalse,
	}
}

func (n Ternary) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

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
