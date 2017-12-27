package expr

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func (n Isset) Name() string {
	return "Isset"
}

type Isset struct {
	name      string
	variables []node.Node
}

func NewIsset(variables []node.Node) node.Node {
	return Isset{
		"Isset",
		variables,
	}
}

func (n Isset) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.variables != nil {
		fmt.Fprintf(out, "\n%vvariables:", indent+"  ")
		for _, nn := range n.variables {
			nn.Print(out, indent+"    ")
		}
	}
}
