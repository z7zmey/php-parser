package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n Interface) Name() string {
	return "Interface"
}

type Interface struct {
	name    string
	token   token.Token
	iName   token.Token
	extends []node.Node
	stmts   []node.Node
}

func NewInterface(token token.Token, name token.Token, extends []node.Node, stmts []node.Node) node.Node {
	return Interface{
		"Interface",
		token,
		name,
		extends,
		stmts,
	}
}

func (n Interface) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.iName.Value)

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
