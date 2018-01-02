package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type InstanceOf struct {
	name       string
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
	class      node.Node
}

func NewInstanceOf(expr node.Node, class node.Node) node.Node {
	return InstanceOf{
		"InstanceOf",
		map[string]interface{}{},
		nil,
		expr,
		class,
	}
}

func (n InstanceOf) Name() string {
	return "InstanceOf"
}

func (n InstanceOf) Attributes() map[string]interface{} {
	return n.attributes
}

func (n InstanceOf) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n InstanceOf) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n InstanceOf) Position() *node.Position {
	return n.position
}

func (n InstanceOf) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n InstanceOf) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	if n.class != nil {
		vv := v.GetChildrenVisitor("class")
		n.class.Walk(vv)
	}

	v.LeaveNode(n)
}
