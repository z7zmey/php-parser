package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n InlineHtml) Name() string {
	return "InlineHtml"
}

type InlineHtml struct {
	name  string
	token token.Token
}

func NewInlineHtml(token token.Token) node.Node {
	return InlineHtml{
		"InlineHtml",
		token,
	}
}

func (n InlineHtml) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)
}
