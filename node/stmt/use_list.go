package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type UseList struct {
	node.SimpleNode
	token   token.Token
	useType node.Node
	uses    []node.Node
}

func NewUseList(token token.Token, useType node.Node, uses []node.Node) node.Node {
	return UseList{
		node.SimpleNode{Name: "UseList", Attributes: make(map[string]string)},
		token,
		useType,
		uses,
	}
}

func (n UseList) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.useType != nil {
		fmt.Fprintf(out, "\n%vtype:", indent+"  ")
		n.useType.Print(out, indent+"    ")
	}

	if n.uses != nil {
		fmt.Fprintf(out, "\n%vuses:", indent+"  ")
		for _, nn := range n.uses {
			nn.Print(out, indent+"    ")
		}
	}
}
