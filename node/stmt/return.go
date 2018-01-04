package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Return struct {
	position *node.Position
	Expr     node.Node
}

func NewReturn(Expr node.Node) node.Node {
	return &Return{
		nil,
		Expr,
	}
}

func (n Return) Attributes() map[string]interface{} {
	return nil
}

func (n Return) Position() *node.Position {
	return n.position
}

func (n Return) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Return) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
