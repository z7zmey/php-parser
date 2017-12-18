package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type Print struct {
	node.SimpleNode
	expr node.Node
}

func NewPrint(expression node.Node) node.Node {
	return Print{
		node.SimpleNode{Name: "Print", Attributes: make(map[string]string)},
		expression,
	}
}

func (n Print) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
