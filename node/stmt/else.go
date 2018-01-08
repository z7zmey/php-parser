package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Else struct {
	position *node.Position
	comments *[]comment.Comment
	Stmt     node.Node
}

func NewElse(Stmt node.Node) *Else {
	return &Else{
		nil,
		nil,
		Stmt,
	}
}

func (n Else) Attributes() map[string]interface{} {
	return nil
}

func (n Else) Position() *node.Position {
	return n.position
}

func (n Else) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Else) Comments() *[]comment.Comment {
	return n.comments
}

func (n Else) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n Else) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Stmt != nil {
		vv := v.GetChildrenVisitor("Stmt")
		n.Stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
