package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Else struct {
	node.SimpleNode
	token token.Token
	stmt  node.Node
}

func NewElse(token token.Token, stmt node.Node) node.Node {
	return Else{
		node.SimpleNode{Name: "Else", Attributes: make(map[string]string)},
		token,
		stmt,
	}
}

func (n Else) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.stmt != nil {
		fmt.Fprintf(out, "\n%vstmt:", indent+"  ")
		n.stmt.Print(out, indent+"    ")
	}
}
