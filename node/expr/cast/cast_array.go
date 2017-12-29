package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type CastArray struct {
	Cast
}

func NewCastArray(expr node.Node) node.Node {
	return CastArray{
		Cast{
			"CastArray",
			map[string]interface{}{},
			expr,
		},
	}
}

func (n CastArray) Name() string {
	return "CastArray"
}

func (n CastArray) Attributes() map[string]interface{} {
	return n.attributes
}

func (n CastArray) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
