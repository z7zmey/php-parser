package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
)

type PropertyList struct {
	node.SimpleNode
	modifiers  node.Node
	properties []node.Node
}

func NewPropertyList(modifiers node.Node, properties []node.Node) node.Node {
	return PropertyList{
		node.SimpleNode{Name: "PropertyList", Attributes: make(map[string]string)},
		modifiers,
		properties,
	}
}

func (n PropertyList) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [- -]", indent, n.Name)

	if n.modifiers != nil {
		fmt.Fprintf(out, "\n%vmodifiers:", indent+"  ")
		n.modifiers.Print(out, indent+"    ")
	}

	if n.properties != nil {
		fmt.Fprintf(out, "\n%vproperties:", indent+"  ")
		for _, nn := range n.properties {
			nn.Print(out, indent+"    ")
		}
	}
}
