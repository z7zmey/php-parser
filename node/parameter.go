package node

import (
	"fmt"
	"io"
)

type Parameter struct {
	SimpleNode
	variableType Node
	variable     Node
	defaultValue Node
	byRef        bool
	variadic     bool
}

func NewParameter(variableType Node, variable Node, defaultValue Node, byRef bool, variadic bool) Node {
	return Parameter{
		SimpleNode{Name: "Parameter", Attributes: make(map[string]string)},
		variableType,
		variable,
		defaultValue,
		byRef,
		variadic,
	}
}

func (n Parameter) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)
	fmt.Fprintf(out, "\n%vbyRef: %t", indent+"  ", n.byRef)
	fmt.Fprintf(out, "\n%vvariadic: %t", indent+"  ", n.variadic)

	if n.variableType != nil {
		fmt.Fprintf(out, "\n%vvariableType:", indent+"  ")
		n.variableType.Print(out, indent+"    ")
	}

	if n.variable != nil {
		fmt.Fprintf(out, "\n%vvariable:", indent+"  ")
		n.variable.Print(out, indent+"    ")
	}

	if n.defaultValue != nil {
		fmt.Fprintf(out, "\n%vdefaultValue:", indent+"  ")
		n.defaultValue.Print(out, indent+"    ")
	}

}
