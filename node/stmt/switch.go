package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Switch struct {
	node.SimpleNode
	token token.Token
	cond  node.Node
	cases []node.Node
}

func NewSwitch(token token.Token, cond node.Node, cases []node.Node) node.Node {
	return Switch{
		node.SimpleNode{Name: "Switch", Attributes: make(map[string]string)},
		token,
		cond,
		cases,
	}
}

func (n Switch) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)

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
