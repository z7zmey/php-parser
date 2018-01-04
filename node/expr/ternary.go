package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Ternary struct {
	attributes map[string]interface{}
	position   *node.Position
	condition  node.Node
	ifTrue     node.Node
	ifFalse    node.Node
}

func NewTernary(condition node.Node, ifTrue node.Node, ifFalse node.Node) node.Node {
	return &Ternary{
		map[string]interface{}{},
		nil,
		condition,
		ifTrue,
		ifFalse,
	}
}

func (n Ternary) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Ternary) Position() *node.Position {
	return n.position
}

func (n Ternary) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Ternary) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.condition != nil {
		vv := v.GetChildrenVisitor("condition")
		n.condition.Walk(vv)
	}

	if n.ifTrue != nil {
		vv := v.GetChildrenVisitor("ifTrue")
		n.ifTrue.Walk(vv)
	}

	if n.ifFalse != nil {
		vv := v.GetChildrenVisitor("ifFalse")
		n.ifFalse.Walk(vv)
	}

	v.LeaveNode(n)
}
