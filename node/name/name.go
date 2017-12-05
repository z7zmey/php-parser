package name

import (
	"fmt"
	"github.com/z7zmey/php-parser/node"
	"io"
)

type Name struct {
	node.SimpleNode
	parts []node.Node
}

func NewName(parts []node.Node) node.Node {
	return Name{
		node.SimpleNode{Name: "Name", Attributes: make(map[string]string)},
		parts,
	}
}

func (n Name) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v", indent, n.Name)
	fmt.Fprintf(out, "\n%vparts:", indent+"  ",)
	for _, nn := range n.parts {
		nn.Print(out, indent+"    ")
	}
}