package cast

import (
	"github.com/z7zmey/php-parser/node"
)

func (n CastInt) Name() string {
	return "CastInt"
}

type CastInt struct {
	Cast
}

func NewCastInt(expr node.Node) node.Node {
	return CastInt{
		Cast{
			"CastInt",
			expr,
		},
	}
}

func (n CastInt) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
