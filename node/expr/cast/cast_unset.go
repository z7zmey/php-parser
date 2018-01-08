package cast

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type CastUnset struct {
	Cast
}

func NewCastUnset(Expr node.Node) *CastUnset {
	return &CastUnset{
		Cast{
			nil,
			nil,
			Expr,
		},
	}
}

func (n CastUnset) Attributes() map[string]interface{} {
	return nil
}

func (n CastUnset) Position() *node.Position {
	return n.position
}

func (n CastUnset) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n CastUnset) Comments() *[]comment.Comment {
	return n.comments
}

func (n CastUnset) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n CastUnset) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
