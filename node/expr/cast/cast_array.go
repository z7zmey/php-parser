package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type CastArray struct {
	Cast
}

func NewCastArray(expr node.Node) node.Node {
	return &CastArray{
		Cast{
			map[string]interface{}{},
			nil,
			expr,
		},
	}
}

func (n CastArray) Attributes() map[string]interface{} {
	return n.attributes
}

func (n CastArray) Position() *node.Position {
	return n.position
}

func (n CastArray) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
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
