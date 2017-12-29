package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type CastBool struct {
	Cast
}

func NewCastBool(expr node.Node) node.Node {
	return CastBool{
		Cast{
			"CastBool",
			map[string]interface{}{},
			expr,
		},
	}
}

func (n CastBool) Name() string {
	return "CastBool"
}

func (n CastBool) Attributes() map[string]interface{} {
	return n.attributes
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
