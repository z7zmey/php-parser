package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n Eval) Name() string {
	return "Eval"
}

type Eval struct {
	name string
	expr node.Node
}

func NewEval(expression node.Node) node.Node {
	return Eval{
		"Eval",
		expression,
	}
}

func (n Eval) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
