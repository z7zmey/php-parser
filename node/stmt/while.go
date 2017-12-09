package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type While struct {
	node.SimpleNode
	token token.Token
	cond  node.Node
	stmt  node.Node
}

func NewWhile(token token.Token, cond node.Node, stmt node.Node) node.Node {
	return While{
		node.SimpleNode{Name: "While", Attributes: make(map[string]string)},
		token,
		cond,
		stmt,
	}
}

func (n While) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.cond != nil {
		fmt.Fprintf(out, "\n%vcond:", indent+"  ")
		n.cond.Print(out, indent+"    ")
	}

	if n.stmt != nil {
		fmt.Fprintf(out, "\n%vstmt:", indent+"  ")
		n.stmt.Print(out, indent+"    ")
	}
}
