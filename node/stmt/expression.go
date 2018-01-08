package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Expression struct {
	position *node.Position
	comments *[]comment.Comment
	Expr     node.Node
}

func NewExpression(Expr node.Node) *Expression {
	return &Expression{
		nil,
		nil,
		Expr,
	}
}

func (n Expression) Attributes() map[string]interface{} {
	return nil
}

func (n Expression) Position() *node.Position {
	return n.position
}

func (n Expression) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Expression) Comments() *[]comment.Comment {
	return n.comments
}

func (n Expression) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n Expression) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
