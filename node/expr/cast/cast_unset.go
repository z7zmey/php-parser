package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type CastUnset struct {
	Cast
}

func NewCastUnset(Expr node.Node) *CastUnset {
	return &CastUnset{
		Cast{
			Expr,
		},
	}
}

func (n *CastUnset) Attributes() map[string]interface{} {
	return nil
}

func (n *CastUnset) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
