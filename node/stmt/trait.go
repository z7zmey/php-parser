package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Trait struct {
	node.SimpleNode
	token token.Token
	stmts []node.Node
}

//TODO: stmts myst be []node.Node
func NewTrait(token token.Token, stmts []node.Node) node.Node {
	return Trait{
		node.SimpleNode{Name: "Trait", Attributes: make(map[string]string)},
		token,
		stmts,
	}
}

func (n Trait) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.stmts != nil {
		fmt.Fprintf(out, "\n%vstmts:", indent+"  ")
		for _, nn := range n.stmts {
			nn.Print(out, indent+"    ")
		}
	}
}
