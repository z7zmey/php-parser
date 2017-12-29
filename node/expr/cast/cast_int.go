package cast

import (
	"github.com/z7zmey/php-parser/node"
)

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

func (n CastInt) Name() string {
	return "CastInt"
}

func (n CastInt) Attributes() map[string]interface{} {
	return nil
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
