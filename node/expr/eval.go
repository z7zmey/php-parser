package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type Eval struct {
	node.SimpleNode
	expr node.Node
}

func NewEval(expression node.Node) node.Node {
	return Eval{
		node.SimpleNode{Name: "Eval", Attributes: make(map[string]string)},
		expression,
	}
}

func (n Eval) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
