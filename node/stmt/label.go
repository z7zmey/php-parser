package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Label struct {
	node.SimpleNode
	token token.Token
}

func NewLabel(token token.Token) node.Node {
	return Label{
		node.SimpleNode{Name: "Label", Attributes: make(map[string]string)},
		token,
	}
}

func (n Label) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)
}
