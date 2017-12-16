package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type IncludeOnce struct {
	node.SimpleNode
	expr node.Node
}

func NewIncludeOnce(expression node.Node) node.Node {
	return IncludeOnce{
		node.SimpleNode{Name: "IncludeOnce", Attributes: make(map[string]string)},
		expression,
	}
}

func (n IncludeOnce) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
