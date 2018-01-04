package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Expression struct {
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
}

func NewExpression(expr node.Node) node.Node {
	return &Expression{
		map[string]interface{}{},
		nil,
		expr,
	}
}

func (n Expression) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Expression) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Expression) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Expression) Position() *node.Position {
	return n.position
}

func (n Expression) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Expression) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
