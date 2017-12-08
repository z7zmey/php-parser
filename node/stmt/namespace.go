package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Namespace struct {
	node.SimpleNode
	token token.Token
	name  node.Node
	stmts []node.Node
}

func NewNamespace(token token.Token, name node.Node, stmts []node.Node) node.Node {
	return Namespace{
		node.SimpleNode{Name: "Namespace", Attributes: make(map[string]string)},
		token,
		name,
		stmts,
	}
}

func (n Namespace) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.name != nil {
		fmt.Fprintf(out, "\n%vname:", indent+"  ")
		n.name.Print(out, indent+"    ")
	}

	if n.stmts != nil {
		fmt.Fprintf(out, "\n%vstmts:", indent+"  ")
		for _, nn := range n.stmts {
			nn.Print(out, indent+"    ")
		}
	}
}
