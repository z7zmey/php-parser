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
			nil,
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

func (n CastString) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n CastString) Position() *node.Position {
	return n.position
}

func (n CastString) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
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
