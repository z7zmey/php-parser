package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type PropertyList struct {
	attributes map[string]interface{}
	position   *node.Position
	modifiers  []node.Node
	properties []node.Node
}

func NewPropertyList(modifiers []node.Node, properties []node.Node) node.Node {
	return &PropertyList{
		map[string]interface{}{},
		nil,
		modifiers,
		properties,
	}
}

func (n PropertyList) Attributes() map[string]interface{} {
	return n.attributes
}

func (n PropertyList) Position() *node.Position {
	return n.position
}

func (n PropertyList) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n PropertyList) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.modifiers != nil {
		vv := v.GetChildrenVisitor("modifiers")
		for _, nn := range n.modifiers {
			nn.Walk(vv)
		}
	}

	if n.properties != nil {
		vv := v.GetChildrenVisitor("properties")
		for _, nn := range n.properties {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
