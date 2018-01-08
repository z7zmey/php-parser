package binary_op

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type BooleanOr struct {
	BinaryOp
}

func NewBooleanOr(Variable node.Node, Expression node.Node) *BooleanOr {
	return &BooleanOr{
		BinaryOp{
			nil,
			nil,
			Variable,
			Expression,
		},
	}
}

func (n BooleanOr) Attributes() map[string]interface{} {
	return nil
}

func (n BooleanOr) Position() *node.Position {
	return n.position
}

func (n BooleanOr) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n BooleanOr) Comments() *[]comment.Comment {
	return n.comments
}

func (n BooleanOr) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n BooleanOr) Walk(v node.Visitor) {
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
