package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type CastUnset struct {
	Cast
}

func NewCastUnset(expr node.Node) node.Node {
	return CastUnset{
		Cast{
			"CastUnset",
			map[string]interface{}{},
			expr,
		},
	}
}

func (n CastUnset) Name() string {
	return "CastUnset"
}

func (n CastUnset) Attributes() map[string]interface{} {
	return n.attributes
}

func (n CastUnset) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
