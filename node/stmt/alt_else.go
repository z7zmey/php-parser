package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type AltElse struct {
	position *node.Position
	comments *[]comment.Comment
	Stmt     node.Node
}

func NewAltElse(Stmt node.Node) node.Node {
	return &AltElse{
		nil,
		nil,
		Stmt,
	}
}

func (n AltElse) Attributes() map[string]interface{} {
	return nil
}

func (n AltElse) Position() *node.Position {
	return n.position
}

func (n AltElse) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n AltElse) Comments() *[]comment.Comment {
	return n.comments
}

func (n AltElse) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n AltElse) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Stmt != nil {
		vv := v.GetChildrenVisitor("Stmt")
		n.Stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
