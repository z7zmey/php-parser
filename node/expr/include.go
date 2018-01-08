package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Include struct {
	position *node.Position
	comments *[]comment.Comment
	Expr     node.Node
}

func NewInclude(Expression node.Node) *Include {
	return &Include{
		nil,
		nil,
		Expression,
	}
}

func (n Include) Attributes() map[string]interface{} {
	return nil
}

func (n Include) Position() *node.Position {
	return n.position
}

func (n Include) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Include) Comments() *[]comment.Comment {
	return n.comments
}

func (n Include) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n Include) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
