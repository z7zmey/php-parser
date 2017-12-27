package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n Else) Name() string {
	return "Else"
}

type Else struct {
	name  string
	token token.Token
	stmt  node.Node
}

func NewElse(token token.Token, stmt node.Node) node.Node {
	return Else{
		"Else",
		token,
		stmt,
	}
}

func (n Else) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.stmt != nil {
		fmt.Fprintf(out, "\n%vstmt:", indent+"  ")
		n.stmt.Print(out, indent+"    ")
	}
}
