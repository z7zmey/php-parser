package binary_op

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type ShiftLeft struct {
	BinaryOp
}

func NewShiftLeft(Variable node.Node, Expression node.Node) *ShiftLeft {
	return &ShiftLeft{
		BinaryOp{
			nil,
			nil,
			Variable,
			Expression,
		},
	}
}

func (n *ShiftLeft) Attributes() map[string]interface{} {
	return nil
}

func (n *ShiftLeft) Position() *node.Position {
	return n.position
}

func (n *ShiftLeft) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *ShiftLeft) Comments() []comment.Comment {
	return n.comments
}

func (n *ShiftLeft) SetComments(c []comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *ShiftLeft) Walk(v node.Visitor) {
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
