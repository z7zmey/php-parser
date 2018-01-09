package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type CastBool struct {
	Cast
}

func NewCastBool(Expr node.Node) *CastBool {
	return &CastBool{
		Cast{
			Expr,
		},
	}
}

func (n *CastBool) Attributes() map[string]interface{} {
	return nil
}

func (n *CastBool) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
