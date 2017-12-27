package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n Case) Name() string {
	return "Case"
}

type Case struct {
	name  string
	token token.Token
	cond  node.Node
	stmts []node.Node
}

func NewCase(token token.Token, cond node.Node, stmts []node.Node) node.Node {
	return Case{
		"Case",
		token,
		cond,
		stmts,
	}
}

func (n Case) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)

	fmt.Fprintf(out, "\n%vcond:", indent+"  ")
	n.cond.Print(out, indent+"    ")

	if n.stmts != nil {
		fmt.Fprintf(out, "\n%vstmts:", indent+"  ")
		for _, nn := range n.stmts {
			nn.Print(out, indent+"    ")
		}
	}
}
