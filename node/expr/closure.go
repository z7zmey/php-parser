package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type Closure struct {
	node.SimpleNode
	params      []node.Node
	uses        []node.Node
	returnType  node.Node
	stmts       []node.Node
	isReturnRef bool
	isStatic    bool
}

func NewClosure(params []node.Node, uses []node.Node, returnType node.Node, stmts []node.Node, isStatic bool, isReturnRef bool) node.Node {
	return Closure{
		node.SimpleNode{Name: "Closure", Attributes: make(map[string]string)},
		params,
		uses,
		returnType,
		stmts,
		isReturnRef,
		isStatic,
	}
}

func (n Closure) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	fmt.Fprintf(out, "\n%vis static: %t", indent+"  ", n.isStatic)
	fmt.Fprintf(out, "\n%vis return ref: %t", indent+"  ", n.isReturnRef)

	if n.params != nil {
		fmt.Fprintf(out, "\n%vparams:", indent+"  ")
		for _, nn := range n.params {
			nn.Print(out, indent+"    ")
		}
	}

	if n.uses != nil {
		fmt.Fprintf(out, "\n%vuses:", indent+"  ")
		for _, nn := range n.uses {
			nn.Print(out, indent+"    ")
		}
	}

	if n.returnType != nil {
		fmt.Fprintf(out, "\n%vreturn type:", indent+"  ")
		n.returnType.Print(out, indent+"    ")
	}

	if n.stmts != nil {
		fmt.Fprintf(out, "\n%vstmts:", indent+"  ")
		for _, nn := range n.stmts {
			nn.Print(out, indent+"    ")
		}
	}
}
