package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type Exit struct {
	node.SimpleNode
	expr  node.Node
	isDie bool
}

func NewExit(expr node.Node, isDie bool) node.Node {
	return Exit{
		node.SimpleNode{Name: "Exit", Attributes: make(map[string]string)},
		expr,
		isDie,
	}
}

func (n Exit) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)
	fmt.Fprintf(out, "\n%vis die: %t", indent+"  ", n.isDie)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
