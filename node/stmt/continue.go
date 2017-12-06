package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Continue struct {
	node.SimpleNode
	token token.Token
	expr node.Node
}

func NewContinue(token token.Token, expr node.Node) node.Node {
	return Continue{
		node.SimpleNode{Name: "Continue", Attributes: make(map[string]string)},
		token,
		expr,
	}
}

func (n Continue) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
