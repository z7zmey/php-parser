package name

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n NamePart) Name() string {
	return "NamePart"
}

type NamePart struct {
	name  string
	token token.Token
}

func NewNamePart(token token.Token) node.Node {
	return NamePart{
		"NamePart",
		token,
	}
}

func (n NamePart) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)
}
