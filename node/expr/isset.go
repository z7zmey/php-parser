package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type Isset struct {
	node.SimpleNode
	variables []node.Node
}

func NewIsset(variables []node.Node) node.Node {
	return Isset{
		node.SimpleNode{Name: "Isset", Attributes: make(map[string]string)},
		variables,
	}
}

func (n Isset) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.variables != nil {
		fmt.Fprintf(out, "\n%vvariables:", indent+"  ")
		for _, nn := range n.variables {
			nn.Print(out, indent+"    ")
		}
	}
}
