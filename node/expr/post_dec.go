package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n PostDec) Name() string {
	return "PostDec"
}

type PostDec struct {
	name     string
	variable node.Node
}

func NewPostDec(variableession node.Node) node.Node {
	return PostDec{
		"PostDec",
		variableession,
	}
}

func (n PostDec) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.variable != nil {
		fmt.Fprintf(out, "\n%vvariable:", indent+"  ")
		n.variable.Print(out, indent+"    ")
	}
}
