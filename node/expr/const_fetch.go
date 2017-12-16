package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type ConstFetch struct {
	node.SimpleNode
	name node.Node
}

func NewConstFetch(name node.Node) node.Node {
	return ConstFetch{
		node.SimpleNode{Name: "ConstFetch", Attributes: make(map[string]string)},
		name,
	}
}

func (n ConstFetch) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.name != nil {
		fmt.Fprintf(out, "\n%vname:", indent+"  ")
		n.name.Print(out, indent+"    ")
	}
}
