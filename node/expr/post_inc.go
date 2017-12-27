package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n PostInc) Name() string {
	return "PostInc"
}

type PostInc struct {
	name     string
	variable node.Node
}

func NewPostInc(variableession node.Node) node.Node {
	return PostInc{
		"PostInc",
		variableession,
	}
}

func (n PostInc) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.variable != nil {
		fmt.Fprintf(out, "\n%vvariable:", indent+"  ")
		n.variable.Print(out, indent+"    ")
	}
}
