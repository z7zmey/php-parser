package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type InstanceOf struct {
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
	Class      node.Node
}

func NewInstanceOf(expr node.Node, Class node.Node) node.Node {
	return &InstanceOf{
		map[string]interface{}{},
		nil,
		expr,
		Class,
	}
}

func (n InstanceOf) Attributes() map[string]interface{} {
	return n.attributes
}

func (n InstanceOf) Position() *node.Position {
	return n.position
}

func (n InstanceOf) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n InstanceOf) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	if n.Class != nil {
		vv := v.GetChildrenVisitor("Class")
		n.Class.Walk(vv)
	}

	v.LeaveNode(n)
}
