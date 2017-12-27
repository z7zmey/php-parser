package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n TraitUse) Name() string {
	return "TraitUse"
}

type TraitUse struct {
	name        string
	token       token.Token
	traits      []node.Node
	adaptations []node.Node
}

//TODO: traits myst be []node.Node
func NewTraitUse(token token.Token, traits []node.Node, adaptations []node.Node) node.Node {
	return TraitUse{
		"TraitUse",
		token,
		traits,
		adaptations,
	}
}

func (n TraitUse) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.traits != nil {
		fmt.Fprintf(out, "\n%vtraits:", indent+"  ")
		for _, nn := range n.traits {
			nn.Print(out, indent+"    ")
		}
	}

	if n.adaptations != nil {
		fmt.Fprintf(out, "\n%vadaptations:", indent+"  ")
		for _, nn := range n.adaptations {
			nn.Print(out, indent+"    ")
		}
	}
}
