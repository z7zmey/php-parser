package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Array struct {
	node.SimpleNode
	opentToken token.Token
	closeToken token.Token
	items []node.Node
	isShortSyntax bool
}

func NewArray(opentToken token.Token, closeToken token.Token, items []node.Node, isShortSyntax bool) node.Node {
	return Array{
		node.SimpleNode{Name: "Array", Attributes: make(map[string]string)},
		opentToken,
		closeToken,
		items,
		isShortSyntax,
	}
}

func (n Array) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d]", indent, n.Name, n.opentToken.StartLine, n.closeToken.EndLine)
	fmt.Fprintf(out, "\n%visShortSyntax: %t", indent+"  ", n.isShortSyntax)

	if n.items != nil {
		fmt.Fprintf(out, "\n%vitems:", indent+"  ")
		for _, nn := range n.items {
			nn.Print(out, indent+"    ")
		}
	}
}
