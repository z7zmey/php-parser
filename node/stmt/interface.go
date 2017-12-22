package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Interface struct {
	node.SimpleNode
	token   token.Token
	name    token.Token
	extends []node.Node
	stmts   []node.Node
}

func NewInterface(token token.Token, name token.Token, extends []node.Node, stmts []node.Node) node.Node {
	return Interface{
		node.SimpleNode{Name: "Interface", Attributes: make(map[string]string)},
		token,
		name,
		extends,
		stmts,
	}
}

func (n Interface) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.name.Value)

	if n.extends != nil {
		fmt.Fprintf(out, "\n%vextends:", indent+"  ")
		for _, nn := range n.extends {
			nn.Print(out, indent+"    ")
		}
	}

	if n.stmts != nil {
		fmt.Fprintf(out, "\n%vstmts:", indent+"  ")
		for _, nn := range n.stmts {
			nn.Print(out, indent+"    ")
		}
	}
}
