package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Case struct {
	node.SimpleNode
	token token.Token
	cond  node.Node
	stmts node.Node
}

//TODO: stmts myst be []node.Node
func NewCase(token token.Token, cond node.Node, stmts node.Node) node.Node {
	return Case{
		node.SimpleNode{Name: "Case", Attributes: make(map[string]string)},
		token,
		cond,
		stmts,
	}
}

func (n Case) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)
	fmt.Fprintf(out, "\n%vcond:", indent+"  ")
	n.cond.Print(out, indent+"    ")
	fmt.Fprintf(out, "\n%vstmts:", indent+"  ")
	n.stmts.Print(out, indent+"    ")
}
