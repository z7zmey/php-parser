package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type CastString struct {
	Cast
}

func NewCastString(Expr node.Node) *CastString {
	return &CastString{
		Cast{
			Expr,
		},
	}
}

func (n *CastString) Attributes() map[string]interface{} {
	return nil
}

func (n *CastString) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
