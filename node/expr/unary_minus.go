package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type UnaryMinus struct {
	position *node.Position
	comments *[]comment.Comment
	Expr     node.Node
}

func NewUnaryMinus(Expression node.Node) node.Node {
	return &UnaryMinus{
		nil,
		nil,
		Expression,
	}
}

func (n UnaryMinus) Attributes() map[string]interface{} {
	return nil
}

func (n UnaryMinus) Position() *node.Position {
	return n.position
}

func (n UnaryMinus) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n UnaryMinus) Comments() *[]comment.Comment {
	return n.comments
}

func (n UnaryMinus) SetComments(c []comment.Comment) node.Node {
	n.comments = &c
	return n
}

func (n UnaryMinus) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
