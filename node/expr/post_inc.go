package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type PostInc struct {
	node.SimpleNode
	variable node.Node
}

func NewPostInc(variableession node.Node) node.Node {
	return PostInc{
		node.SimpleNode{Name: "PostInc", Attributes: make(map[string]string)},
		variableession,
	}
}

func (n PostInc) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.variable != nil {
		fmt.Fprintf(out, "\n%vvariable:", indent+"  ")
		n.variable.Print(out, indent+"    ")
	}
}
