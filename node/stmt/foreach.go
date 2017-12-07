package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Foreach struct {
	node.SimpleNode
	token    token.Token
	expr     node.Node
	key      node.Node
	variable node.Node
	stmt     node.Node
}

func NewForeach(token token.Token, expr node.Node, key node.Node, variable node.Node, stmt node.Node) node.Node {
	return Foreach{
		node.SimpleNode{Name: "Foreach", Attributes: make(map[string]string)},
		token,
		expr,
		key,
		variable,
		stmt,
	}
}

func (n Foreach) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.expr != nil {
		fmt.Fprintf(out, "\n%vexpr:", indent+"  ")
		n.expr.Print(out, indent+"    ")
	}

	if n.key != nil {
		fmt.Fprintf(out, "\n%vkey:", indent+"  ")
		n.key.Print(out, indent+"    ")
	}

	if n.variable != nil {
		fmt.Fprintf(out, "\n%vvariable:", indent+"  ")
		n.variable.Print(out, indent+"    ")
	}

	if n.stmt != nil {
		fmt.Fprintf(out, "\n%vstmt:", indent+"  ")
		n.stmt.Print(out, indent+"    ")
	}
}
