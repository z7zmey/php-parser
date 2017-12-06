package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type ElseIf struct {
	node.SimpleNode
	token token.Token
	cond  node.Node
	stmts []node.Node
}

func NewElseIf(token token.Token, cond node.Node, stmts []node.Node) node.Node {
	return ElseIf{
		node.SimpleNode{Name: "ElseIf", Attributes: make(map[string]string)},
		token,
		cond,
		stmts,
	}
}

func (n ElseIf) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.cond != nil {
		fmt.Fprintf(out, "\n%vcond:", indent+"  ")
		n.cond.Print(out, indent+"    ")
	}

	if n.stmts != nil {
		fmt.Fprintf(out, "\n%vstmts:", indent+"  ")
		for _, nn := range n.stmts {
			nn.Print(out, indent+"    ")
		}
	}
}
