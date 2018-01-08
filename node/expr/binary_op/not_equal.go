package binary_op

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type NotEqual struct {
	BinaryOp
}

func NewNotEqual(Variable node.Node, Expression node.Node) *NotEqual {
	return &NotEqual{
		BinaryOp{
			nil,
			nil,
			Variable,
			Expression,
		},
	}
}

func (n *NotEqual) Attributes() map[string]interface{} {
	return nil
}

func (n *NotEqual) Position() *node.Position {
	return n.position
}

func (n *NotEqual) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *NotEqual) Comments() *[]comment.Comment {
	return n.comments
}

func (n *NotEqual) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *NotEqual) Walk(v node.Visitor) {
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
