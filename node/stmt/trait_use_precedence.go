package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func(n TraitUsePrecedence) Name() string {
	return "TraitUsePrecedence"
}

type TraitUsePrecedence struct {
	name      string
	ref       node.Node
	insteadof node.Node
}

func NewTraitUsePrecedence(ref node.Node, insteadof node.Node) node.Node {
	return TraitUsePrecedence{
		"TraitUsePrecedence",
		ref,
		insteadof,
	}
}

func (n TraitUsePrecedence) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.ref != nil {
		fmt.Fprintf(out, "\n%vmethod", indent+"  ")
		n.ref.Print(out, indent+"    ")
	}

	if n.insteadof != nil {
		fmt.Fprintf(out, "\n%vinsteadof:", indent+"  ")
		n.insteadof.Print(out, indent+"    ")
	}
}
