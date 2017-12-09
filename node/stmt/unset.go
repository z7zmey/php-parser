package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Unset struct {
	node.SimpleNode
	token token.Token
	vars  []node.Node
}

func NewUnset(token token.Token, vars []node.Node) node.Node {
	return Unset{
		node.SimpleNode{Name: "Unset", Attributes: make(map[string]string)},
		token,
		vars,
	}
}

func (n Unset) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.vars != nil {
		fmt.Fprintf(out, "\n%vvars:", indent+"  ")
		for _, nn := range n.vars {
			nn.Print(out, indent+"    ")
		}
	}
}
