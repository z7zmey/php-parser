package node

import (
	"fmt"
	"io"
)

type Argument struct {
	SimpleNode
	expr     Node
	variadic bool
}

func NewArgument(expression Node, variadic bool) Node {
	return Argument{
		SimpleNode{Name: "Argument", Attributes: make(map[string]string)},
		expression,
		variadic,
	}
}

func (n Argument) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)
	fmt.Fprintf(out, "\n%vvariadic: %t", indent+"  ", n.variadic)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
