package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

func(n PropertyList) Name() string {
	return "PropertyList"
}

type PropertyList struct {
	name       string
	modifiers  []node.Node
	properties []node.Node
}

func NewPropertyList(modifiers []node.Node, properties []node.Node) node.Node {
	return PropertyList{
		"PropertyList",
		modifiers,
		properties,
	}
}

func (n PropertyList) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.name)

	if n.modifiers != nil {
		fmt.Fprintf(out, "\n%vmodifiers:", indent+"  ")
		for _, nn := range n.modifiers {
			nn.Print(out, indent+"    ")
		}
	}

	if n.properties != nil {
		fmt.Fprintf(out, "\n%vproperties:", indent+"  ")
		for _, nn := range n.properties {
			nn.Print(out, indent+"    ")
		}
	}
}
