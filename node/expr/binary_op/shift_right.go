package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type ShiftRight struct {
	BinaryOp
}

func NewShiftRight(Variable node.Node, Expression node.Node) *ShiftRight {
	return &ShiftRight{
		BinaryOp{
			nil,
			Variable,
			Expression,
		},
	}
}

func (n *ShiftRight) Attributes() map[string]interface{} {
	return nil
}

func (n *ShiftRight) Position() *node.Position {
	return n.position
}

func (n *ShiftRight) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *ShiftRight) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Left != nil {
		vv := v.GetChildrenVisitor("Left")
		n.Left.Walk(vv)
	}

	if n.Right != nil {
		vv := v.GetChildrenVisitor("Right")
		n.Right.Walk(vv)
	}

	v.LeaveNode(n)
}
