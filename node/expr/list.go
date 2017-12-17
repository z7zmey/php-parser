package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type List struct {
	node.SimpleNode
	items []node.Node
}

func NewList(items []node.Node) node.Node {
	return List{
		node.SimpleNode{Name: "List", Attributes: make(map[string]string)},
		items,
	}
}

func (n List) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.items != nil {
		fmt.Fprintf(out, "\n%vitems:", indent+"  ")
		for _, nn := range n.items {
			nn.Print(out, indent+"    ")
		}
	}
}
