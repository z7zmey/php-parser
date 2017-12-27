package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n Yield) Name() string {
	return "Yield"
}

type Yield struct {
	name  string
	key   node.Node
	value node.Node
}

func NewYield(key node.Node, value node.Node) node.Node {
	return Yield{
		"Yield",
		key,
		value,
	}
}

func (n Yield) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.key != nil {
		fmt.Fprintf(out, "\n%vkey:", indent+"  ")
		n.key.Print(out, indent+"    ")
	}

	if n.value != nil {
		fmt.Fprintf(out, "\n%vvalue:", indent+"  ")
		n.value.Print(out, indent+"    ")
	}
}
