package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type TraitUsePrecedence struct {
	node.SimpleNode
	ref       node.Node
	insteadof node.Node
}

func NewTraitUsePrecedence(ref node.Node, insteadof node.Node) node.Node {
	return TraitUsePrecedence{
		node.SimpleNode{Name: "TraitUsePrecedence", Attributes: make(map[string]string)},
		ref,
		insteadof,
	}
}

func (n TraitUsePrecedence) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.ref != nil {
		fmt.Fprintf(out, "\n%vmethod", indent+"  ")
		n.ref.Print(out, indent+"    ")
	}

	if n.insteadof != nil {
		fmt.Fprintf(out, "\n%vinsteadof:", indent+"  ")
		n.insteadof.Print(out, indent+"    ")
	}
}
