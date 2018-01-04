package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Empty struct {
	position *node.Position
	Expr     node.Node
}

func NewEmpty(Expression node.Node) node.Node {
	return &Empty{
		nil,
		Expression,
	}
}

func (n Empty) Attributes() map[string]interface{} {
	return nil
}

func (n Empty) Position() *node.Position {
	return n.position
}

func (n Empty) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Empty) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
