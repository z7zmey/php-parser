package cast

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type CastDouble struct {
	Cast
}

func NewCastDouble(Expr node.Node) *CastDouble {
	return &CastDouble{
		Cast{
			nil,
			nil,
			Expr,
		},
	}
}

func (n *CastDouble) Attributes() map[string]interface{} {
	return nil
}

func (n *CastDouble) Position() *node.Position {
	return n.position
}

func (n *CastDouble) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *CastDouble) Comments() []comment.Comment {
	return n.comments
}

func (n *CastDouble) SetComments(c []comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *CastDouble) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
