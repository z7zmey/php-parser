package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type ConstList struct {
	node.SimpleNode
	token  token.Token
	consts []node.Node
}

func NewConstList(token token.Token, consts []node.Node) node.Node {
	return ConstList{
		node.SimpleNode{Name: "ConstList", Attributes: make(map[string]string)},
		token,
		consts,
	}
}

func (n ConstList) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.consts != nil {
		fmt.Fprintf(out, "\n%vconsts:", indent+"  ")
		for _, nn := range n.consts {
			nn.Print(out, indent+"    ")
		}
	}
}
