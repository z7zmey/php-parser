package cast

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type CastBool struct {
	Cast
}

func NewCastBool(Expr node.Node) *CastBool {
	return &CastBool{
		Cast{
			nil,
			nil,
			Expr,
		},
	}
}

func (n *CastBool) Attributes() map[string]interface{} {
	return nil
}

func (n *CastBool) Position() *node.Position {
	return n.position
}

func (n *CastBool) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *CastBool) Comments() []comment.Comment {
	return n.comments
}

func (n *CastBool) SetComments(c []comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *CastBool) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
