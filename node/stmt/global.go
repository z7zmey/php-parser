package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n Global) Name() string {
	return "Global"
}

type Global struct {
	name  string
	token token.Token
	vars  []node.Node
}

func NewGlobal(token token.Token, vars []node.Node) node.Node {
	return Global{
		"Global",
		token,
		vars,
	}
}

func (n Global) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.vars != nil {
		fmt.Fprintf(out, "\n%vvars:", indent+"  ")
		for _, nn := range n.vars {
			nn.Print(out, indent+"    ")
		}
	}
}
