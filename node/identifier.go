package node

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/token"
)

type Identifier struct {
	name  string
	token token.Token
}

func (n Identifier) Name() string {
	return "Identifier"
}

func NewIdentifier(token token.Token) Node {
	return Identifier{
		"Identifier",
		token,
	}
}

func (n Identifier) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)
	fmt.Fprintf(out, "\n%vname: %q", indent+"  ", n.token.Value)
}
