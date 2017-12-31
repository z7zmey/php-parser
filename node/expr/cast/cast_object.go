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
			map[string]interface{}{},
			nil,
			expr,
		},
	}
}

func (n CastObject) Name() string {
	return "CastObject"
}

func (n CastObject) Attributes() map[string]interface{} {
	return n.attributes
}

func (n CastObject) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n CastObject) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n CastObject) Position() *node.Position {
	return n.position
}

func (n CastObject) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
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
