package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n Goto) Name() string {
	return "Goto"
}

type Goto struct {
	name  string
	token token.Token
	label token.Token
}

func NewGoto(token token.Token, name token.Token) node.Node {
	return Goto{
		"Goto",
		token,
		name,
	}
}

func (n Goto) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.label.EndLine, n.label.Value)
}
