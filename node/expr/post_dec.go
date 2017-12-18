package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type PostDec struct {
	node.SimpleNode
	variable node.Node
}

func NewPostDec(variableession node.Node) node.Node {
	return PostDec{
		node.SimpleNode{Name: "PostDec", Attributes: make(map[string]string)},
		variableession,
	}
}

func (n PostDec) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.variable != nil {
		fmt.Fprintf(out, "\n%vvariable:", indent+"  ")
		n.variable.Print(out, indent+"    ")
	}
}
