package cast

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type CastString struct {
	Cast
}

func NewCastString(Expr node.Node) *CastString {
	return &CastString{
		Cast{
			nil,
			nil,
			Expr,
		},
	}
}

func (n *CastString) Attributes() map[string]interface{} {
	return nil
}

func (n *CastString) Position() *node.Position {
	return n.position
}

func (n *CastString) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *CastString) Comments() []comment.Comment {
	return n.comments
}

func (n *CastString) SetComments(c []comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *CastString) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
