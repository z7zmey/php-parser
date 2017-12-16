package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type ClusureUse struct {
	node.SimpleNode
	variable node.Node
	byRef    bool
}

func NewClusureUse(variable node.Node, byRef bool) node.Node {
	return ClusureUse{
		node.SimpleNode{Name: "ClusureUse", Attributes: make(map[string]string)},
		variable,
		byRef,
	}
}

func (n ClusureUse) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)
	fmt.Fprintf(out, "\n%vby ref: %t", indent+"  ", n.byRef)

	if n.variable != nil {
		fmt.Fprintf(out, "\n%vvariable:", indent+"  ")
		n.variable.Print(out, indent+"    ")
	}
}
