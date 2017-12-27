package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n For) Name() string {
	return "For"
}

type For struct {
	name  string
	token token.Token
	init  []node.Node
	cond  []node.Node
	loop  []node.Node
	stmt  node.Node
}

func NewFor(token token.Token, init []node.Node, cond []node.Node, loop []node.Node, stmt node.Node) node.Node {
	return For{
		"For",
		token,
		init,
		cond,
		loop,
		stmt,
	}
}

func (n For) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.init != nil {
		fmt.Fprintf(out, "\n%vinit:", indent+"  ")
		for _, nn := range n.init {
			nn.Print(out, indent+"    ")
		}
	}

	if n.cond != nil {
		fmt.Fprintf(out, "\n%vcond:", indent+"  ")
		for _, nn := range n.cond {
			nn.Print(out, indent+"    ")
		}
	}

	if n.loop != nil {
		fmt.Fprintf(out, "\n%vloop:", indent+"  ")
		for _, nn := range n.loop {
			nn.Print(out, indent+"    ")
		}
	}

	if n.stmt != nil {
		fmt.Fprintf(out, "\n%vstmt:", indent+"  ")
		n.stmt.Print(out, indent+"    ")
	}
}
