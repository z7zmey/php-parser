package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n Echo) Name() string {
	return "Echo"
}

type Echo struct {
	name  string
	token token.Token
	exprs []node.Node
}

func NewEcho(token token.Token, exprs []node.Node) node.Node {
	return Echo{
		"Echo",
		token,
		exprs,
	}
}

func (n Echo) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.exprs != nil {
		fmt.Fprintf(out, "\n%vexprs:", indent+"  ")
		for _, nn := range n.exprs {
			nn.Print(out, indent+"    ")
		}
	}
}
