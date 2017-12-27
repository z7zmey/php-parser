package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n ClusureUse) Name() string {
	return "ClusureUse"
}

type ClusureUse struct {
	name     string
	variable node.Node
	byRef    bool
}

func NewClusureUse(variable node.Node, byRef bool) node.Node {
	return ClusureUse{
		"ClusureUse",
		variable,
		byRef,
	}
}

func (n ClusureUse) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)
	fmt.Fprintf(out, "\n%vby ref: %t", indent+"  ", n.byRef)

	if n.variable != nil {
		fmt.Fprintf(out, "\n%vvariable:", indent+"  ")
		n.variable.Print(out, indent+"    ")
	}
}
