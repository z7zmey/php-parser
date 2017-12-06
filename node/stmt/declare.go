package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Declare struct {
	node.SimpleNode
	token  token.Token
	consts []node.Node
	stmts  []node.Node
}

func NewDeclare(token token.Token, consts []node.Node, stmts []node.Node) node.Node {
	return Declare{
		node.SimpleNode{Name: "Declare", Attributes: make(map[string]string)},
		token,
		consts,
		stmts,
	}
}

func (n Declare) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.consts != nil {
		fmt.Fprintf(out, "\n%vconsts:", indent+"  ")
		for _, nn := range n.consts {
			nn.Print(out, indent+"    ")
		}
	}

	if n.stmts != nil {
		fmt.Fprintf(out, "\n%vstmts:", indent+"  ")
		for _, nn := range n.stmts {
			nn.Print(out, indent+"    ")
		}
	}
}
