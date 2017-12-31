package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type CastDouble struct {
	Cast
}

func NewCastDouble(expr node.Node) node.Node {
	return CastDouble{
		Cast{
			"CastDouble",
			map[string]interface{}{},
			expr,
		},
	}
}

func (n CastDouble) Name() string {
	return "CastDouble"
}

func (n CastDouble) Attributes() map[string]interface{} {
	return n.attributes
}

func (n CastDouble) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n CastDouble) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n CastDouble) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
