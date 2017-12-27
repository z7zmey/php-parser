package scalar

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n Encapsed) Name() string {
	return "Encapsed"
}

type Encapsed struct {
	name       string
	startToken token.Token
	endToken   token.Token
	parts      []node.Node
}

func NewEncapsed(startToken token.Token, parts []node.Node, endToken token.Token) node.Node {
	return Encapsed{
		"Encapsed",
		startToken,
		endToken,
		parts,
	}
}

func (n Encapsed) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d]", indent, n.name, n.startToken.StartLine, n.endToken.EndLine)

	if n.parts != nil {
		fmt.Fprintf(out, "\n%vparts:", indent+"  ")
		for _, nn := range n.parts {
			nn.Print(out, indent+"    ")
		}
	}
}
