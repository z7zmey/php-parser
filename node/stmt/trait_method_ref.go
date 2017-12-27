package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n TraitMethodRef) Name() string {
	return "TraitMethodRef"
}

type TraitMethodRef struct {
	name   string
	trait  node.Node
	method token.Token
}

func NewTraitMethodRef(trait node.Node, method token.Token) node.Node {
	return TraitMethodRef{
		"TraitMethodRef",
		trait,
		method,
	}
}

func (n TraitMethodRef) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -] %q", indent, n.name, n.method.Value)

	if n.trait != nil {
		fmt.Fprintf(out, "\n%vtrait", indent+"  ")
		n.trait.Print(out, indent+"    ")
	}
}
