package binary_op

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type SmallerOrEqual struct {
	BinaryOp
}

func NewSmallerOrEqual(Variable node.Node, Expression node.Node) node.Node {
	return &SmallerOrEqual{
		BinaryOp{
			nil,
			nil,
			Variable,
			Expression,
		},
	}
}

func (n SmallerOrEqual) Attributes() map[string]interface{} {
	return nil
}

func (n SmallerOrEqual) Position() *node.Position {
	return n.position
}

func (n SmallerOrEqual) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n SmallerOrEqual) Comments() *[]comment.Comment {
	return n.comments
}

func (n SmallerOrEqual) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n SmallerOrEqual) Walk(v node.Visitor) {
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
