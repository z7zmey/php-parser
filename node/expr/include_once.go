package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type IncludeOnce struct {
	position *node.Position
	comments *[]comment.Comment
	Expr     node.Node
}

func NewIncludeOnce(Expression node.Node) node.Node {
	return &IncludeOnce{
		nil,
		nil,
		Expression,
	}
}

func (n IncludeOnce) Attributes() map[string]interface{} {
	return nil
}

func (n IncludeOnce) Position() *node.Position {
	return n.position
}

func (n IncludeOnce) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n IncludeOnce) Comments() *[]comment.Comment {
	return n.comments
}

func (n IncludeOnce) SetComments(c []comment.Comment) node.Node {
	n.comments = &c
	return n
}

func (n IncludeOnce) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
