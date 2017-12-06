package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Finally struct {
	node.SimpleNode
	token token.Token
	stmts node.Node
}

//TODO: stmts myst be []node.Node
func NewFinally(token token.Token, stmts node.Node) node.Node {
	return Finally{
		node.SimpleNode{Name: "Finally", Attributes: make(map[string]string)},
		token,
		stmts,
	}
}

func (n Finally) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)

	fmt.Fprintf(out, "\n%vstmts:", indent+"  ")
	n.stmts.Print(out, indent+"    ")
}
