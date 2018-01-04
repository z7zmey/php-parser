package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type PropertyFetch struct {
	attributes map[string]interface{}
	position   *node.Position
	variable   node.Node
	property   node.Node
}

func NewPropertyFetch(variable node.Node, property node.Node) node.Node {
	return &PropertyFetch{
		map[string]interface{}{},
		nil,
		variable,
		property,
	}
}

func (n PropertyFetch) Attributes() map[string]interface{} {
	return n.attributes
}

func (n PropertyFetch) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n PropertyFetch) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n PropertyFetch) Position() *node.Position {
	return n.position
}

func (n PropertyFetch) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n PropertyFetch) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	if n.property != nil {
		vv := v.GetChildrenVisitor("property")
		n.property.Walk(vv)
	}

	v.LeaveNode(n)
}
