package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type StaticVar struct {
	node.SimpleNode
	token token.Token
	expr  node.Node
}

func NewStaticVar(token token.Token, expr node.Node) node.Node {
	return StaticVar{
		node.SimpleNode{Name: "StaticVar", Attributes: make(map[string]string)},
		token,
		expr,
	}
}

func (n StaticVar) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
