package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Require struct {
	position *node.Position
	Expr     node.Node
}

func NewRequire(Expression node.Node) *Require {
	return &Require{
		nil,
		Expression,
	}
}

func (n *Require) Attributes() map[string]interface{} {
	return nil
}

func (n *Require) Position() *node.Position {
	return n.position
}

func (n *Require) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *Require) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
