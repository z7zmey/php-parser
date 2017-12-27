package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n ClassConstList) Name() string {
	return "ClassConstList"
}

type ClassConstList struct {
	name      string
	token     token.Token
	modifiers []node.Node
	consts    []node.Node
}

func NewClassConstList(token token.Token, modifiers []node.Node, consts []node.Node) node.Node {
	return ClassConstList{
		"ClassConstList",
		token,
		modifiers,
		consts,
	}
}

func (n ClassConstList) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.modifiers != nil {
		fmt.Fprintf(out, "\n%vmotifiers:", indent+"  ")
		for _, nn := range n.modifiers {
			nn.Print(out, indent+"    ")
		}
	}

	if n.consts != nil {
		fmt.Fprintf(out, "\n%vconsts:", indent+"  ")
		for _, nn := range n.consts {
			nn.Print(out, indent+"    ")
		}
	}
}
