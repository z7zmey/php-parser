package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type CastArray struct {
	Cast
}

func NewCastArray(Expr node.Node) *CastArray {
	return &CastArray{
		Cast{
			Expr,
		},
	}
}

func (n *CastArray) Attributes() map[string]interface{} {
	return nil
}

func (n *CastArray) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
