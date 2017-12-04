package scalar

import (
	"fmt"
	"github.com/z7zmey/php-parser/token"
	"github.com/z7zmey/php-parser/node"
	"io"
)


type Encapsed struct {
	node.SimpleNode
	startToken token.Token
	endToken token.Token
	parts []node.Node
}


func NewEncapsed(startToken token.Token, parts []node.Node, endToken token.Token) node.Node {
	return Encapsed{
		node.SimpleNode{Name: "Encapsed", Attributes: make(map[string]string)},
		startToken,
		endToken,
		parts,
	}
}

func (n Encapsed) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d]", indent, n.Name, n.startToken.StartLine, n.endToken.EndLine)
	fmt.Fprintf(out, "\n%vparts:", indent+"  ",)
	for _, nn := range n.parts {
		nn.Print(out, indent+"    ")
	}
}
