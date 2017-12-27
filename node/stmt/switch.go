package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n Switch) Name() string {
	return "Switch"
}

type Switch struct {
	name  string
	token token.Token
	cond  node.Node
	cases []node.Node
}

func NewSwitch(token token.Token, cond node.Node, cases []node.Node) node.Node {
	return Switch{
		"Switch",
		token,
		cond,
		cases,
	}
}

func (n Switch) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.cond != nil {
		fmt.Fprintf(out, "\n%vcond:", indent+"  ")
		n.cond.Print(out, indent+"    ")
	}

	if n.cases != nil {
		fmt.Fprintf(out, "\n%vcases:", indent+"  ")
		for _, nn := range n.cases {
			nn.Print(out, indent+"    ")
		}
	}
}
