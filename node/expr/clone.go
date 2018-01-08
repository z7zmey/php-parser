package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Clone struct {
	position *node.Position
	Expr     node.Node
}

func NewClone(Expression node.Node) *Clone {
	return &Clone{
		nil,
		Expression,
	}
}

func (n *Clone) Attributes() map[string]interface{} {
	return nil
}

func (n *Clone) Position() *node.Position {
	return n.position
}

func (n *Clone) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *Clone) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
