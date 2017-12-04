package scalar

import (
	"fmt"
	"github.com/z7zmey/php-parser/token"
	"github.com/z7zmey/php-parser/node"
	"io"
)


type MagicConstant struct {
	node.SimpleNode
	token token.Token
}


func NewMagicConstant(token token.Token) node.Node {
	return String{
		node.SimpleNode{Name: "MagicConstant", Attributes: make(map[string]string)},
		token,
	}
}

func (n MagicConstant) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)
	for _, nn := range n.Children {
		nn.Print(out, indent+"  ")
	}
}
