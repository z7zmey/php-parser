package binary_op

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Spaceship struct {
	BinaryOp
}

func NewSpaceship(Variable node.Node, Expression node.Node) *Spaceship {
	return &Spaceship{
		BinaryOp{
			nil,
			nil,
			Variable,
			Expression,
		},
	}
}

func (n Spaceship) Attributes() map[string]interface{} {
	return nil
}

func (n Spaceship) Position() *node.Position {
	return n.position
}

func (n Spaceship) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Spaceship) Comments() *[]comment.Comment {
	return n.comments
}

func (n Spaceship) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n Spaceship) Walk(v node.Visitor) {
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
