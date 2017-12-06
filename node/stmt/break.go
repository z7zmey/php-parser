package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Break struct {
	node.SimpleNode
	token token.Token
	expr  node.Node
}

func NewBreak(token token.Token, expr node.Node) node.Node {
	return Break{
		node.SimpleNode{Name: "Break", Attributes: make(map[string]string)},
		token,
		expr,
	}
}

func (n Break) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)
	fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
	n.expr.Print(out, indent+"    ")
}
