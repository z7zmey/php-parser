package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Continue struct {
	position *node.Position
	comments []comment.Comment
	Expr     node.Node
}

func NewContinue(Expr node.Node) *Continue {
	return &Continue{
		nil,
		nil,
		Expr,
	}
}

func (n *Continue) Attributes() map[string]interface{} {
	return nil
}

func (n *Continue) Position() *node.Position {
	return n.position
}

func (n *Continue) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *Continue) Comments() []comment.Comment {
	return n.comments
}

func (n *Continue) SetComments(c []comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *Continue) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
