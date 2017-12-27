package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n PreInc) Name() string {
	return "PreInc"
}

type PreInc struct {
	name     string
	variable node.Node
}

func NewPreInc(variableession node.Node) node.Node {
	return PreInc{
		"PreInc",
		variableession,
	}
}

func (n PreInc) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.variable != nil {
		fmt.Fprintf(out, "\n%vvariable:", indent+"  ")
		n.variable.Print(out, indent+"    ")
	}
}
