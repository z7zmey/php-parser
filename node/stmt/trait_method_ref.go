package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type TraitMethodRef struct {
	node.SimpleNode
	trait  node.Node
	method token.Token
}

func NewTraitMethodRef(trait node.Node, method token.Token) node.Node {
	return TraitMethodRef{
		node.SimpleNode{Name: "TraitMethodRef", Attributes: make(map[string]string)},
		trait,
		method,
	}
}

func (n TraitMethodRef) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -] %q", indent, n.Name, n.method.Value)

	if n.trait != nil {
		fmt.Fprintf(out, "\n%vtrait", indent+"  ")
		n.trait.Print(out, indent+"    ")
	}
}
