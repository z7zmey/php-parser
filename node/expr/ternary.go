package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Ternary struct {
	position  *node.Position
	Condition node.Node
	IfTrue    node.Node
	IfFalse   node.Node
}

func NewTernary(Condition node.Node, IfTrue node.Node, IfFalse node.Node) node.Node {
	return &Ternary{
		nil,
		Condition,
		IfTrue,
		IfFalse,
	}
}

func (n Ternary) Attributes() map[string]interface{} {
	return nil
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

	if n.Condition != nil {
		vv := v.GetChildrenVisitor("Condition")
		n.Condition.Walk(vv)
	}

	if n.IfTrue != nil {
		vv := v.GetChildrenVisitor("IfTrue")
		n.IfTrue.Walk(vv)
	}

	if n.IfFalse != nil {
		vv := v.GetChildrenVisitor("IfFalse")
		n.IfFalse.Walk(vv)
	}

	v.LeaveNode(n)
}
