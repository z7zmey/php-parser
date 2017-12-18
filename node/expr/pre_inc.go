package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type PreInc struct {
	node.SimpleNode
	variable node.Node
}

func NewPreInc(variableession node.Node) node.Node {
	return PreInc{
		node.SimpleNode{Name: "PreInc", Attributes: make(map[string]string)},
		variableession,
	}
}

func (n PreInc) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.variable != nil {
		fmt.Fprintf(out, "\n%vvariable:", indent+"  ")
		n.variable.Print(out, indent+"    ")
	}
}
