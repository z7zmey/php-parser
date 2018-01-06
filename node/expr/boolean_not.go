package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type BooleanNot struct {
	position *node.Position
	comments *[]comment.Comment
	Expr     node.Node
}

func NewBooleanNot(Expression node.Node) node.Node {
	return &BooleanNot{
		nil,
		nil,
		Expression,
	}
}

func (n BooleanNot) Attributes() map[string]interface{} {
	return nil
}

func (n BooleanNot) Position() *node.Position {
	return n.position
}

func (n BooleanNot) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n BooleanNot) Comments() *[]comment.Comment {
	return n.comments
}

func (n BooleanNot) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n BooleanNot) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
