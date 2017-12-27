package scalar

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n MagicConstant) Name() string {
	return "MagicConstant"
}

type MagicConstant struct {
	name  string
	token token.Token
}

func NewMagicConstant(token token.Token) node.Node {
	return MagicConstant{
		"MagicConstant",
		token,
	}
}

func (n MagicConstant) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)
}
