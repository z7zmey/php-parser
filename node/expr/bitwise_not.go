package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type BitwiseNot struct {
	name       string
	attributes map[string]interface{}
	position *node.Position
	expr       node.Node
}

func NewBitwiseNot(expression node.Node) node.Node {
	return BitwiseNot{
		"BitwiseNot",
		map[string]interface{}{},
		nil,
		expression,
	}
}

func (n BitwiseNot) Name() string {
	return "BitwiseNot"
}

func (n BitwiseNot) Attributes() map[string]interface{} {
	return n.attributes
}

func (n BitwiseNot) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n BitwiseNot) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n BitwiseNot) Position() *node.Position {
	return n.position
}

func (n BitwiseNot) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n BitwiseNot) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
