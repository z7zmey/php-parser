package node

import (
	"fmt"
	"io"
)

type Nullable struct {
	SimpleNode
	expr Node
}

func NewNullable(expression Node) Node {
	return Nullable{
		SimpleNode{Name: "Nullable", Attributes: make(map[string]string)},
		expression,
	}
}

func (n Nullable) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
