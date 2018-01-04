package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type UnaryMinus struct {
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
}

func NewUnaryMinus(expression node.Node) node.Node {
	return &UnaryMinus{
		map[string]interface{}{},
		nil,
		expression,
	}
}

func (n UnaryMinus) Attributes() map[string]interface{} {
	return n.attributes
}

func (n UnaryMinus) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n UnaryMinus) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n UnaryMinus) Position() *node.Position {
	return n.position
}

func (n UnaryMinus) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n UnaryMinus) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
