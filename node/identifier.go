package node

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/token"
)

type Identifier struct {
	SimpleNode
	name token.Token
}

func NewIdentifier(name token.Token) Node {
	return Identifier{
		SimpleNode{Name: "Identifier", Attributes: make(map[string]string)},
		name,
	}
}

func (n Identifier) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)
	fmt.Fprintf(out, "\n%vname: %q", indent+"  ", n.name.Value)
}
