package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type InlineHtml struct {
	node.SimpleNode
	token token.Token
}

func NewInlineHtml(token token.Token) node.Node {
	return InlineHtml{
		node.SimpleNode{Name: "InlineHtml", Attributes: make(map[string]string)},
		token,
	}
}

func (n InlineHtml) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)
}
