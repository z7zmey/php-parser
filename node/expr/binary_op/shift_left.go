package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type ShiftLeft struct {
	BinaryOp
}

func NewShiftLeft(variable node.Node, expression node.Node) node.Node {
	return ShiftLeft{
		BinaryOp{
			"BinaryShiftLeft",
			map[string]interface{}{},
			nil,
			variable,
			expression,
		},
	}
}

func (n ShiftLeft) Name() string {
	return "ShiftLeft"
}

func (n ShiftLeft) Attributes() map[string]interface{} {
	return n.attributes
}

func (n ShiftLeft) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n ShiftLeft) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n ShiftLeft) Position() *node.Position {
	return n.position
}

func (n ShiftLeft) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n ShiftLeft) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.left != nil {
		vv := v.GetChildrenVisitor("left")
		n.left.Walk(vv)
	}

	if n.right != nil {
		vv := v.GetChildrenVisitor("right")
		n.right.Walk(vv)
	}

	v.LeaveNode(n)
}
