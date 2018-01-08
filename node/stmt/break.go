package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Break struct {
	position *node.Position
	Expr     node.Node
}

func NewBreak(Expr node.Node) *Break {
	return &Break{
		nil,
		Expr,
	}
}

func (n *Break) Attributes() map[string]interface{} {
	return nil
}

func (n *Break) Position() *node.Position {
	return n.position
}

func (n *Break) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *Break) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
