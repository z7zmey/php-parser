package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type StaticPropertyFetch struct {
	attributes map[string]interface{}
	position   *node.Position
	Class      node.Node
	Property   node.Node
}

func NewStaticPropertyFetch(Class node.Node, Property node.Node) node.Node {
	return &StaticPropertyFetch{
		map[string]interface{}{},
		nil,
		Class,
		Property,
	}
}

func (n StaticPropertyFetch) Attributes() map[string]interface{} {
	return n.attributes
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

	if n.Class != nil {
		vv := v.GetChildrenVisitor("Class")
		n.Class.Walk(vv)
	}

	if n.Property != nil {
		vv := v.GetChildrenVisitor("Property")
		n.Property.Walk(vv)
	}

	v.LeaveNode(n)
}
