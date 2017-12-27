package node

import (
	"fmt"
	"io"
)

type Argument struct {
	name     string
	expr     Node
	variadic bool
}

func (n Argument) Name() string {
	return "Argument"
}

func NewArgument(expression Node, variadic bool) Node {
	return Argument{
		"Argument",
		expression,
		variadic,
	}
}

func (n Argument) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)
	fmt.Fprintf(out, "\n%vvariadic: %t", indent+"  ", n.variadic)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}
}
