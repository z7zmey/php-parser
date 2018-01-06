package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type PreDec struct {
	position *node.Position
	comments *[]comment.Comment
	Variable node.Node
}

func NewPreDec(Variable node.Node) node.Node {
	return &PreDec{
		nil,
		nil,
		Variable,
	}
}

func (n PreDec) Attributes() map[string]interface{} {
	return nil
}

func (n PreDec) Position() *node.Position {
	return n.position
}

func (n PreDec) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n PreDec) Comments() *[]comment.Comment {
	return n.comments
}

func (n PreDec) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n PreDec) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	v.LeaveNode(n)
}
