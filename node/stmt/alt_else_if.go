package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type AltElseIf struct {
	position *node.Position
	comments *[]comment.Comment
	Cond     node.Node
	Stmt     node.Node
}

func NewAltElseIf(Cond node.Node, Stmt node.Node) node.Node {
	return &AltElseIf{
		nil,
		nil,
		Cond,
		Stmt,
	}
}

func (n AltElseIf) Attributes() map[string]interface{} {
	return nil
}

func (n AltElseIf) Position() *node.Position {
	return n.position
}

func (n AltElseIf) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n AltElseIf) Comments() *[]comment.Comment {
	return n.comments
}

func (n AltElseIf) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n AltElseIf) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Cond != nil {
		vv := v.GetChildrenVisitor("Cond")
		n.Cond.Walk(vv)
	}

	if n.Stmt != nil {
		vv := v.GetChildrenVisitor("Stmt")
		n.Stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
