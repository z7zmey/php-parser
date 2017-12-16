package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type Empty struct {
	node.SimpleNode
	expr node.Node
}

func NewEmpty(expression node.Node) node.Node {
	return Empty{
		node.SimpleNode{Name: "Empty", Attributes: make(map[string]string)},
		expression,
	}
}

func (n Empty) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
