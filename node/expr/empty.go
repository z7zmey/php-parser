package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n Empty) Name() string {
	return "Empty"
}

type Empty struct {
	name string
	expr node.Node
}

func NewEmpty(expression node.Node) node.Node {
	return Empty{
		"Empty",
		expression,
	}
}

func (n Empty) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
