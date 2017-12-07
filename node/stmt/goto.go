package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Goto struct {
	node.SimpleNode
	token token.Token
	name  token.Token
}

func NewGoto(token token.Token, name token.Token) node.Node {
	return Goto{
		node.SimpleNode{Name: "Goto", Attributes: make(map[string]string)},
		token,
		name,
	}
}

func (n Goto) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.name.EndLine, n.name.Value)
}
