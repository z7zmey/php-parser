package scalar

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n Dnumber) Name() string {
	return "Dnumber"
}

type Dnumber struct {
	name  string
	token token.Token
}

func NewDnumber(token token.Token) node.Node {
	return Dnumber{
		"Dnumber",
		token,
	}
}

func (n Dnumber) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)
}
