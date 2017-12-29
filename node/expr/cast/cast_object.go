package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type CastObject struct {
	Cast
}

func NewCastObject(expr node.Node) node.Node {
	return CastObject{
		Cast{
			"CastObject",
			expr,
		},
	}
}

func (n CastObject) Name() string {
	return "CastObject"
}

func (n CastObject) Attributes() map[string]interface{} {
	return nil
}

func (n CastObject) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
