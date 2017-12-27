package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n Declare) Name() string {
	return "Declare"
}

type Declare struct {
	name   string
	token  token.Token
	consts []node.Node
	stmt   node.Node
}

func NewDeclare(token token.Token, consts []node.Node, stmt node.Node) node.Node {
	return Declare{
		"Declare",
		token,
		consts,
		stmt,
	}
}

func (n Declare) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.consts != nil {
		fmt.Fprintf(out, "\n%vconsts:", indent+"  ")
		for _, nn := range n.consts {
			nn.Print(out, indent+"    ")
		}
	}

	if n.stmt != nil {
		fmt.Fprintf(out, "\n%vstmt:", indent+"  ")
		n.stmt.Print(out, indent+"    ")
	}
}
