package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n ClassMethod) Name() string {
	return "ClassMethod"
}

type ClassMethod struct {
	name        string
	token       token.Token
	modifiers   []node.Node
	isReturnRef bool
	params      []node.Node
	returnType  node.Node
	stmts       []node.Node
}

func NewClassMethod(token token.Token, modifiers []node.Node, isReturnRef bool, params []node.Node, returnType node.Node, stmts []node.Node) node.Node {
	return ClassMethod{
		"ClassMethod",
		token,
		modifiers,
		isReturnRef,
		params,
		returnType,
		stmts,
	}
}

func (n ClassMethod) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.modifiers != nil {
		fmt.Fprintf(out, "\n%vmodifiers:", indent+"  ")
		for _, nn := range n.modifiers {
			nn.Print(out, indent+"    ")
		}
	}

	fmt.Fprintf(out, "\n%vreturn ref: %t", indent+"  ", n.isReturnRef)

	if n.params != nil {
		fmt.Fprintf(out, "\n%vparams:", indent+"  ")
		for _, nn := range n.params {
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
