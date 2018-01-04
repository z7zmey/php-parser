package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type CastDouble struct {
	Cast
}

func NewCastDouble(Expr node.Node) node.Node {
	return &CastDouble{
		Cast{
			nil,
			Expr,
		},
	}
}

func (n CastDouble) Attributes() map[string]interface{} {
	return nil
}

func (n CastDouble) Position() *node.Position {
	return n.position
}

func (n CastDouble) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n CastDouble) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
