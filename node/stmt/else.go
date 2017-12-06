package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Else struct {
	node.SimpleNode
	token token.Token
	stmts []node.Node
}

func NewElse(token token.Token, stmts []node.Node) node.Node {
	return Else{
		node.SimpleNode{Name: "Else", Attributes: make(map[string]string)},
		token,
		stmts,
	}
}

func (n Else) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.stmts != nil {
		fmt.Fprintf(out, "\n%vstmts:", indent+"  ")
		for _, nn := range n.stmts {
			nn.Print(out, indent+"    ")
		}
	}
}
