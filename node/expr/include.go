package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type Include struct {
	node.SimpleNode
	expr node.Node
}

func NewInclude(expression node.Node) node.Node {
	return Include{
		node.SimpleNode{Name: "Include", Attributes: make(map[string]string)},
		expression,
	}
}

func (n Include) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
