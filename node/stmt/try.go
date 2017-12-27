package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n Try) Name() string {
	return "Try"
}

type Try struct {
	name    string
	token   token.Token
	stmts   []node.Node
	catches []node.Node
	finally node.Node
}

func NewTry(token token.Token, stmts []node.Node, catches []node.Node, finally node.Node) node.Node {
	return Try{
		"Try",
		token,
		stmts,
		catches,
		finally,
	}
}

func (n Try) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.stmts != nil {
		fmt.Fprintf(out, "\n%vstmts:", indent+"  ")
		for _, nn := range n.stmts {
			nn.Print(out, indent+"    ")
		}
	}

	if n.catches != nil {
		fmt.Fprintf(out, "\n%vcatches:", indent+"  ")
		for _, nn := range n.catches {
			nn.Print(out, indent+"    ")
		}
	}

	if n.finally != nil {
		fmt.Fprintf(out, "\n%vfinally:", indent+"  ")
		n.finally.Print(out, indent+"    ")
	}
}
