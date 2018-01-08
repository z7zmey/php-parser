package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type UnaryPlus struct {
	position *node.Position
	comments []comment.Comment
	Expr     node.Node
}

func NewUnaryPlus(Expression node.Node) *UnaryPlus {
	return &UnaryPlus{
		nil,
		nil,
		Expression,
	}
}

func (n *UnaryPlus) Attributes() map[string]interface{} {
	return nil
}

func (n *UnaryPlus) Position() *node.Position {
	return n.position
}

func (n *UnaryPlus) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *UnaryPlus) Comments() []comment.Comment {
	return n.comments
}

func (n *UnaryPlus) SetComments(c []comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *UnaryPlus) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
