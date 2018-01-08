package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type NotIdentical struct {
	BinaryOp
}

func NewNotIdentical(Variable node.Node, Expression node.Node) *NotIdentical {
	return &NotIdentical{
		BinaryOp{
			nil,
			Variable,
			Expression,
		},
	}
}

func (n *NotIdentical) Attributes() map[string]interface{} {
	return nil
}

func (n *NotIdentical) Position() *node.Position {
	return n.position
}

func (n *NotIdentical) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *NotIdentical) Walk(v node.Visitor) {
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
