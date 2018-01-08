package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Continue struct {
	position *node.Position
	Expr     node.Node
}

func NewContinue(Expr node.Node) *Continue {
	return &Continue{
		nil,
		Expr,
	}
}

func (n *Continue) Attributes() map[string]interface{} {
	return nil
}

func (n *Continue) Position() *node.Position {
	return n.position
}

func (n *Continue) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *Continue) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
