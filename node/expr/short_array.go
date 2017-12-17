package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type ShortArray struct {
	node.SimpleNode
	opentToken token.Token
	closeToken token.Token
	items      []node.Node
}

func NewShortArray(opentToken token.Token, closeToken token.Token, items []node.Node) node.Node {
	return ShortArray{
		node.SimpleNode{Name: "ShortArray", Attributes: make(map[string]string)},
		opentToken,
		closeToken,
		items,
	}
}

func (n ShortArray) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d]", indent, n.Name, n.opentToken.StartLine, n.closeToken.EndLine)

	if n.items != nil {
		fmt.Fprintf(out, "\n%vitems:", indent+"  ")
		for _, nn := range n.items {
			nn.Print(out, indent+"    ")
		}
	}
}
