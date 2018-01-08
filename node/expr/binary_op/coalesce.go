package binary_op

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Coalesce struct {
	BinaryOp
}

func NewCoalesce(Variable node.Node, Expression node.Node) *Coalesce {
	return &Coalesce{
		BinaryOp{
			nil,
			nil,
			Variable,
			Expression,
		},
	}
}

func (n Coalesce) Attributes() map[string]interface{} {
	return nil
}

func (n Coalesce) Position() *node.Position {
	return n.position
}

func (n Coalesce) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Coalesce) Comments() *[]comment.Comment {
	return n.comments
}

func (n Coalesce) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n Coalesce) Walk(v node.Visitor) {
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
