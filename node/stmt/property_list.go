package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

func (n PropertyList) Name() string {
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

func (n PropertyList) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.modifiers != nil {
		vv := v.Children("modifiers")
		for _, nn := range n.modifiers {
			nn.Walk(vv)
		}
	}

	if n.properties != nil {
		vv := v.Children("properties")
		for _, nn := range n.properties {
			nn.Walk(vv)
		}
	}
}
