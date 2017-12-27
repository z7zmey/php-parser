package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n Exit) Name() string {
	return "Exit"
}

type Exit struct {
	name  string
	expr  node.Node
	isDie bool
}

func NewExit(expr node.Node, isDie bool) node.Node {
	return Exit{
		"Exit",
		expr,
		isDie,
	}
}

func (n Exit) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)
	fmt.Fprintf(out, "\n%vis die: %t", indent+"  ", n.isDie)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
