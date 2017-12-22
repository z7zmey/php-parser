package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type ErrorSuppress struct {
	node.SimpleNode
	expr node.Node
}

func NewErrorSuppress(expression node.Node) node.Node {
	return ErrorSuppress{
		node.SimpleNode{Name: "ErrorSuppress", Attributes: make(map[string]string)},
		expression,
	}
}

func (n ErrorSuppress) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
