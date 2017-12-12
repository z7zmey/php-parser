package cast

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type Cast struct {
	node.SimpleNode
	expr node.Node
}

func NewCast(expr node.Node) node.Node {
	return Cast{
		node.SimpleNode{Name: "Cast", Attributes: make(map[string]string)},
		expr,
	}
}

func (n Cast) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
