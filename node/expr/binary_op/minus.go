package binary_op

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Minus struct {
	BinaryOp
}

func NewMinus(Variable node.Node, Expression node.Node) *Minus {
	return &Minus{
		BinaryOp{
			nil,
			nil,
			Variable,
			Expression,
		},
	}
}

func (n *Minus) Attributes() map[string]interface{} {
	return nil
}

func (n *Minus) Position() *node.Position {
	return n.position
}

func (n *Minus) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *Minus) Comments() []comment.Comment {
	return n.comments
}

func (n *Minus) SetComments(c []comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *Minus) Walk(v node.Visitor) {
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
