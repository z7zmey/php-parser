package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type CastInt struct {
	Cast
}

func NewCastInt(Expr node.Node) node.Node {
	return &CastInt{
		Cast{
			nil,
			Expr,
		},
	}
}

func (n CastInt) Attributes() map[string]interface{} {
	return nil
}

func (n CastInt) Position() *node.Position {
	return n.position
}

func (n CastInt) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n CastInt) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
