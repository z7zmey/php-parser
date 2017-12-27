package name

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n NameNode) Name() string {
	return "Name"
}

type NameNode struct {
	name  string
	parts []node.Node
}

func NewName(parts []node.Node) node.Node {
	return NameNode{
		"Name",
		parts,
	}
}

func (n NameNode) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v", indent, n.name)

	if n.parts != nil {
		fmt.Fprintf(out, "\n%vparts:", indent+"  ")
		for _, nn := range n.parts {
			nn.Print(out, indent+"    ")
		}
	}
}
