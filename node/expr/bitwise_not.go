package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type BitwiseNot struct {
	position *node.Position
	Expr     node.Node
}

func NewBitwiseNot(Expression node.Node) *BitwiseNot {
	return &BitwiseNot{
		nil,
		Expression,
	}
}

func (n *BitwiseNot) Attributes() map[string]interface{} {
	return nil
}

func (n *BitwiseNot) Position() *node.Position {
	return n.position
}

func (n *BitwiseNot) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *BitwiseNot) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
