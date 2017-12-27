package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n Foreach) Name() string {
	return "Foreach"
}

type Foreach struct {
	name     string
	token    token.Token
	expr     node.Node
	key      node.Node
	variable node.Node
	stmt     node.Node
	byRef    bool
}

func NewForeach(token token.Token, expr node.Node, key node.Node, variable node.Node, stmt node.Node, byRef bool) node.Node {
	return Foreach{
		"Foreach",
		token,
		expr,
		key,
		variable,
		stmt,
		byRef,
	}
}

func (n Foreach) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}

	if n.key != nil {
		fmt.Fprintf(out, "\n%vkey:", indent+"  ")
		n.key.Print(out, indent+"    ")
	}

	if n.variable != nil {
		fmt.Fprintf(out, "\n%vvariable[byRef: %t]:", indent+"  ", n.byRef)
		n.variable.Print(out, indent+"    ")
	}

	if n.stmt != nil {
		fmt.Fprintf(out, "\n%vstmt:", indent+"  ")
		n.stmt.Print(out, indent+"    ")
	}
}
