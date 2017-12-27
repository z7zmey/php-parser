package scalar

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n String) Name() string {
	return "String"
}

type String struct {
	name  string
	token token.Token
}

func NewString(token token.Token) node.Node {
	return String{
		"String",
		token,
	}
}

func (n String) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)
}
