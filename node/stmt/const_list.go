package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n ConstList) Name() string {
	return "ConstList"
}

type ConstList struct {
	name   string
	token  token.Token
	consts []node.Node
}

func NewConstList(token token.Token, consts []node.Node) node.Node {
	return ConstList{
		"ConstList",
		token,
		consts,
	}
}

func (n ConstList) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.consts != nil {
		fmt.Fprintf(out, "\n%vconsts:", indent+"  ")
		for _, nn := range n.consts {
			nn.Print(out, indent+"    ")
		}
	}
}
