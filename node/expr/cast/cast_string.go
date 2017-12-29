package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type CastString struct {
	Cast
}

func NewCastString(expr node.Node) node.Node {
	return CastString{
		Cast{
			"CastString",
			expr,
		},
	}
}

func (n CastString) Name() string {
	return "CastString"
}

func (n CastString) Attributes() map[string]interface{} {
	return nil
}

func (n CastString) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
