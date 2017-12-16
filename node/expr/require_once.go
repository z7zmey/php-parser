package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type RequireOnce struct {
	node.SimpleNode
	expr node.Node
}

func NewRequireOnce(expression node.Node) node.Node {
	return RequireOnce{
		node.SimpleNode{Name: "RequireOnce", Attributes: make(map[string]string)},
		expression,
	}
}

func (n RequireOnce) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
