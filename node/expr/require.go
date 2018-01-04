package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Require struct {
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
}

func NewRequire(expression node.Node) node.Node {
	return Require{
		map[string]interface{}{},
		nil,
		expression,
	}
}

func (n Require) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Require) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Require) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Require) Position() *node.Position {
	return n.position
}

func (n Require) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Require) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
