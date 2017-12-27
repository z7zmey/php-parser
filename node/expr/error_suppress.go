package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n ErrorSuppress) Name() string {
	return "ErrorSuppress"
}

type ErrorSuppress struct {
	name string
	expr node.Node
}

func NewErrorSuppress(expression node.Node) node.Node {
	return ErrorSuppress{
		"ErrorSuppress",
		expression,
	}
}

func (n ErrorSuppress) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
