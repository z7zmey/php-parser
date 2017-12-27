package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n List) Name() string {
	return "List"
}

type List struct {
	name  string
	items []node.Node
}

func NewList(items []node.Node) node.Node {
	return List{
		"List",
		items,
	}
}

func (n List) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.items != nil {
		fmt.Fprintf(out, "\n%vitems:", indent+"  ")
		for _, nn := range n.items {
			nn.Print(out, indent+"    ")
		}
	}
}
