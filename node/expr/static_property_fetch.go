package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type StaticPropertyFetch struct {
	attributes map[string]interface{}
	position   *node.Position
	class      node.Node
	property   node.Node
}

func NewStaticPropertyFetch(class node.Node, property node.Node) node.Node {
	return StaticPropertyFetch{
		map[string]interface{}{},
		nil,
		class,
		property,
	}
}

func (n StaticPropertyFetch) Attributes() map[string]interface{} {
	return n.attributes
}

func (n StaticPropertyFetch) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n StaticPropertyFetch) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n StaticPropertyFetch) Position() *node.Position {
	return n.position
}

func (n StaticPropertyFetch) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n StaticPropertyFetch) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.class != nil {
		vv := v.GetChildrenVisitor("class")
		n.class.Walk(vv)
	}

	if n.property != nil {
		vv := v.GetChildrenVisitor("property")
		n.property.Walk(vv)
	}

	v.LeaveNode(n)
}
