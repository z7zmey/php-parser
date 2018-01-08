package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Exit struct {
	position *node.Position
	comments *[]comment.Comment
	Expr     node.Node
	IsDie    bool
}

func NewExit(Expr node.Node, IsDie bool) *Exit {
	return &Exit{
		nil,
		nil,
		Expr,
		IsDie,
	}
}

func (n *Exit) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"IsDie": n.IsDie,
	}
}

func (n *Exit) Position() *node.Position {
	return n.position
}

func (n *Exit) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *Exit) Comments() *[]comment.Comment {
	return n.comments
}

func (n *Exit) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *Exit) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
