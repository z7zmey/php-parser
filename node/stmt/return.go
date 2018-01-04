package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Return struct {
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
}

func NewReturn(expr node.Node) node.Node {
	return &Return{
		map[string]interface{}{},
		nil,
		expr,
	}
}

func (n Return) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Return) Position() *node.Position {
	return n.position
}

func (n Return) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Return) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
