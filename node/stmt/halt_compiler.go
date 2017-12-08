package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type HaltCompiler struct {
	node.SimpleNode
	token token.Token
}

func NewHaltCompiler(token token.Token) node.Node {
	return HaltCompiler{
		node.SimpleNode{Name: "HaltCompiler", Attributes: make(map[string]string)},
		token,
	}
}

func (n HaltCompiler) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)
}
