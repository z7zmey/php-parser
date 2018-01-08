package assign_op

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type ShiftRight struct {
	AssignOp
}

func NewShiftRight(Variable node.Node, Expression node.Node) *ShiftRight {
	return &ShiftRight{
		AssignOp{
			nil,
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

func (n *ShiftRight) Comments() *[]comment.Comment {
	return n.comments
}

func (n *ShiftRight) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *ShiftRight) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	if n.Expression != nil {
		vv := v.GetChildrenVisitor("Expression")
		n.Expression.Walk(vv)
	}

	v.LeaveNode(n)
}
