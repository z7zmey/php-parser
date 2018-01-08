package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type PreInc struct {
	position *node.Position
	comments *[]comment.Comment
	Variable node.Node
}

func NewPreInc(Variable node.Node) *PreInc {
	return &PreInc{
		nil,
		nil,
		Variable,
	}
}

func (n PreInc) Attributes() map[string]interface{} {
	return nil
}

func (n PreInc) Position() *node.Position {
	return n.position
}

func (n PreInc) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n PreInc) Comments() *[]comment.Comment {
	return n.comments
}

func (n PreInc) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n PreInc) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	v.LeaveNode(n)
}
