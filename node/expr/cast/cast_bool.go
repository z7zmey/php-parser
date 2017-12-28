package cast

import (
	"github.com/z7zmey/php-parser/node"
)

func (n CastBool) Name() string {
	return "CastBool"
}

type CastBool struct {
	Cast
}

func NewCastBool(expr node.Node) node.Node {
	return CastBool{
		Cast{
			"CastBool",
			expr,
		},
	}
}

func (n CastBool) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
