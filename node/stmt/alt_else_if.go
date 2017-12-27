package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n AltElseIf) Name() string {
	return "AltElseIf"
}

type AltElseIf struct {
	name  string
	token token.Token
	cond  node.Node
	stmt  node.Node
}

func NewAltElseIf(token token.Token, cond node.Node, stmt node.Node) node.Node {
	return AltElseIf{
		"AltElseIf",
		token,
		cond,
		stmt,
	}
}

func (n AltElseIf) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.cond != nil {
		fmt.Fprintf(out, "\n%vcond:", indent+"  ")
		n.cond.Print(out, indent+"    ")
	}

	if n.stmt != nil {
		fmt.Fprintf(out, "\n%vstmt:", indent+"  ")
		n.stmt.Print(out, indent+"    ")
	}
}
