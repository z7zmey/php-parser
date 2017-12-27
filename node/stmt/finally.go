package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n Finally) Name() string {
	return "Finally"
}

type Finally struct {
	name  string
	token token.Token
	stmts []node.Node
}

func NewFinally(token token.Token, stmts []node.Node) node.Node {
	return Finally{
		"Finally",
		token,
		stmts,
	}
}

func (n Finally) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.stmts != nil {
		fmt.Fprintf(out, "\n%vstmts:", indent+"  ")
		for _, nn := range n.stmts {
			nn.Print(out, indent+"    ")
		}
	}
}
