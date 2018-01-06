package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Eval struct {
	position *node.Position
	comments *[]comment.Comment
	Expr     node.Node
}

func NewEval(Expression node.Node) node.Node {
	return &Eval{
		nil,
		nil,
		Expression,
	}
}

func (n Eval) Attributes() map[string]interface{} {
	return nil
}

func (n Eval) Position() *node.Position {
	return n.position
}

func (n Eval) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Eval) Comments() *[]comment.Comment {
	return n.comments
}

func (n Eval) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n Eval) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
