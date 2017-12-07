package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Do struct {
	node.SimpleNode
	token token.Token
	stmt  node.Node
	cond  node.Node
}

func NewDo(token token.Token, stmt node.Node, cond node.Node) node.Node {
	return Do{
		node.SimpleNode{Name: "Do", Attributes: make(map[string]string)},
		token,
		stmt,
		cond,
	}
}

func (n Do) Print(out io.Writer, indent string) {
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
