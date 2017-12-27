package cast

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type Cast struct {
	name string
	expr node.Node
}

func (n Cast) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
