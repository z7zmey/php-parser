package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type UnaryMinus struct {
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
}

func NewUnaryMinus(Expression node.Node) node.Node {
	return &UnaryMinus{
		map[string]interface{}{},
		nil,
		Expression,
	}
}

func (n UnaryMinus) Attributes() map[string]interface{} {
	return n.attributes
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
