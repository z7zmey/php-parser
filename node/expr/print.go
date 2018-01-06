package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Print struct {
	position *node.Position
	comments *[]comment.Comment
	Expr     node.Node
}

func NewPrint(Expression node.Node) node.Node {
	return &Print{
		nil,
		nil,
		Expression,
	}
}

func (n Print) Attributes() map[string]interface{} {
	return nil
}

func (n Print) Position() *node.Position {
	return n.position
}

func (n Print) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Print) Comments() *[]comment.Comment {
	return n.comments
}

func (n Print) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n Print) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
