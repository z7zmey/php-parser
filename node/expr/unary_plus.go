package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type UnaryPlus struct {
	name       string
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
}

func NewUnaryPlus(expression node.Node) node.Node {
	return UnaryPlus{
		"UnaryPlus",
		map[string]interface{}{},
		nil,
		expression,
	}
}

func (n UnaryPlus) Name() string {
	return "UnaryPlus"
}

func (n UnaryPlus) Attributes() map[string]interface{} {
	return n.attributes
}

func (n UnaryPlus) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n UnaryPlus) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n UnaryPlus) Position() *node.Position {
	return n.position
}

func (n UnaryPlus) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n UnaryPlus) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
