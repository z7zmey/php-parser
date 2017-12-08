package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Nop struct {
	node.SimpleNode
	token token.Token
}

func NewNop(token token.Token) node.Node {
	return Nop{
		node.SimpleNode{Name: "Nop", Attributes: make(map[string]string)},
		token,
	}
}

func (n Nop) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)
}
