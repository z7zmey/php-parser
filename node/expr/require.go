package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Require struct {
	position *node.Position
	comments *[]comment.Comment
	Expr     node.Node
}

func NewRequire(Expression node.Node) node.Node {
	return &Require{
		nil,
		nil,
		Expression,
	}
}

func (n Require) Attributes() map[string]interface{} {
	return nil
}

func (n Require) Position() *node.Position {
	return n.position
}

func (n Require) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Require) Comments() *[]comment.Comment {
	return n.comments
}

func (n Require) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n Require) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
