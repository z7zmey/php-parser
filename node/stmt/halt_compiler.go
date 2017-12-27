package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n HaltCompiler) Name() string {
	return "HaltCompiler"
}

type HaltCompiler struct {
	name  string
	token token.Token
}

func NewHaltCompiler(token token.Token) node.Node {
	return HaltCompiler{
		"HaltCompiler",
		token,
	}
}

func (n HaltCompiler) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)
}
