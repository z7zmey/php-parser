package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type UnaryPlus struct {
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
}

func NewUnaryPlus(Expression node.Node) node.Node {
	return &UnaryPlus{
		map[string]interface{}{},
		nil,
		Expression,
	}
}

func (n UnaryPlus) Attributes() map[string]interface{} {
	return n.attributes
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
