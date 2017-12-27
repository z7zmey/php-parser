package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n Label) Name() string {
	return "Label"
}

type Label struct {
	name  string
	token token.Token
}

func NewLabel(token token.Token) node.Node {
	return Label{
		"Label",
		token,
	}
}

func (n Label) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)
}
