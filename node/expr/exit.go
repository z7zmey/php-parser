package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Exit struct {
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
}

func NewExit(expr node.Node, isDie bool) node.Node {
	return &Exit{
		map[string]interface{}{
			"isDie": isDie,
		},
		nil,
		expr,
	}
}

func (n Exit) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Exit) Position() *node.Position {
	return n.position
}

func (n Exit) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Exit) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
