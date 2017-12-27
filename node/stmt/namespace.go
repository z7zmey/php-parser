package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n Namespace) Name() string {
	return "Namespace"
}

type Namespace struct {
	name  string
	token token.Token
	nName node.Node
	stmts []node.Node
}

func NewNamespace(token token.Token, name node.Node, stmts []node.Node) node.Node {
	return Namespace{
		"Namespace",
		token,
		name,
		stmts,
	}
}

func (n Namespace) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.nName != nil {
		fmt.Fprintf(out, "\n%vname:", indent+"  ")
		n.nName.Print(out, indent+"    ")
	}

	if n.stmts != nil {
		fmt.Fprintf(out, "\n%vstmts:", indent+"  ")
		for _, nn := range n.stmts {
			nn.Print(out, indent+"    ")
		}
	}
}
