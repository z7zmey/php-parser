package binary_op

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type GreaterOrEqual struct {
	BinaryOp
}

func NewGreaterOrEqual(Variable node.Node, Expression node.Node) node.Node {
	return &GreaterOrEqual{
		BinaryOp{
			nil,
			nil,
			Variable,
			Expression,
		},
	}
}

func (n GreaterOrEqual) Attributes() map[string]interface{} {
	return nil
}

func (n GreaterOrEqual) Position() *node.Position {
	return n.position
}

func (n GreaterOrEqual) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n GreaterOrEqual) Comments() *[]comment.Comment {
	return n.comments
}

func (n GreaterOrEqual) SetComments(c []comment.Comment) node.Node {
	n.comments = &c
	return n
}

func (n GreaterOrEqual) Walk(v node.Visitor) {
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
