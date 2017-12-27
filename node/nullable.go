package node

import (
	"fmt"
	"io"
)

type Nullable struct {
	name string
	expr Node
}

func (n Nullable) Name() string {
	return "Nullable"
}

func NewNullable(expression Node) Node {
	return Nullable{
		"Nullable",
		expression,
	}
}

func (n Nullable) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
