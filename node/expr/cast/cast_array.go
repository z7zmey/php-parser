package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type CastArray struct {
	Cast
}

func NewCastArray(Expr node.Node) node.Node {
	return &CastArray{
		Cast{
			nil,
			Expr,
		},
	}
}

func (n CastArray) Attributes() map[string]interface{} {
	return nil
}

func (n CastArray) Position() *node.Position {
	return n.position
}

func (n CastArray) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n CastArray) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
