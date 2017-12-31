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
			map[string]interface{}{},
			expr,
		},
	}
}

func (n CastString) Name() string {
	return "CastString"
}

func (n CastString) Attributes() map[string]interface{} {
	return n.attributes
}

func (n CastString) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n CastString) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
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
