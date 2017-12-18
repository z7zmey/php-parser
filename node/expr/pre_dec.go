package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type PreDec struct {
	node.SimpleNode
	variable node.Node
}

func NewPreDec(variableession node.Node) node.Node {
	return PreDec{
		node.SimpleNode{Name: "PreDec", Attributes: make(map[string]string)},
		variableession,
	}
}

func (n PreDec) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.variable != nil {
		fmt.Fprintf(out, "\n%vvariable:", indent+"  ")
		n.variable.Print(out, indent+"    ")
	}
}
