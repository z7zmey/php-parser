package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n PreDec) Name() string {
	return "PreDec"
}

type PreDec struct {
	name     string
	variable node.Node
}

func NewPreDec(variableession node.Node) node.Node {
	return PreDec{
		"PreDec",
		variableession,
	}
}

func (n PreDec) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.variable != nil {
		fmt.Fprintf(out, "\n%vvariable:", indent+"  ")
		n.variable.Print(out, indent+"    ")
	}
}
