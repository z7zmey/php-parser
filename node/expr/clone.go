package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Clone struct {
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
}

func NewClone(expression node.Node) node.Node {
	return &Clone{
		map[string]interface{}{},
		nil,
		expression,
	}
}

func (n Clone) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Clone) Position() *node.Position {
	return n.position
}

func (n Clone) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Clone) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
