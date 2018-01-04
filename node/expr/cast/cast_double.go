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
			map[string]interface{}{},
			nil,
			expr,
		},
	}
}

func (n CastDouble) Attributes() map[string]interface{} {
	return n.attributes
}

func (n CastDouble) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n CastDouble) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n CastDouble) Position() *node.Position {
	return n.position
}

func (n CastDouble) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
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
