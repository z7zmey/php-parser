package cast

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type CastObject struct {
	Cast
}

func NewCastObject(Expr node.Node) node.Node {
	return &CastObject{
		Cast{
			nil,
			nil,
			Expr,
		},
	}
}

func (n CastObject) Attributes() map[string]interface{} {
	return nil
}

func (n CastObject) Position() *node.Position {
	return n.position
}

func (n CastObject) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n CastObject) Comments() *[]comment.Comment {
	return n.comments
}

func (n CastObject) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n CastObject) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
