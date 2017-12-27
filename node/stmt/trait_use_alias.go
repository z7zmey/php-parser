package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n TraitUseAlias) Name() string {
	return "TraitUseAlias"
}

type TraitUseAlias struct {
	name     string
	ref      node.Node
	modifier node.Node
	alias    token.TokenInterface
}

func NewTraitUseAlias(ref node.Node, modifier node.Node, alias token.TokenInterface) node.Node {
	return TraitUseAlias{
		"TraitUseAlias",
		ref,
		modifier,
		alias,
	}
}

func (n TraitUseAlias) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.alias != nil {
		fmt.Fprintf(out, "\n%valias: %q", indent+"  ", n.alias.GetValue())
	}

	if n.ref != nil {
		fmt.Fprintf(out, "\n%vmethod", indent+"  ")
		n.ref.Print(out, indent+"    ")
	}

	if n.modifier != nil {
		fmt.Fprintf(out, "\n%vmodifier:", indent+"  ")
		n.modifier.Print(out, indent+"    ")
	}
}
