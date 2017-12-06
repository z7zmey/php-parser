package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Variable struct {
	node.SimpleNode
	token token.Token
}

func NewVariable(token token.Token) node.Node {
	return Variable{
		node.SimpleNode{Name: "Variable", Attributes: make(map[string]string)},
		token,
	}
}

func (n Variable) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)
}
