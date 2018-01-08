package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type PostInc struct {
	position *node.Position
	comments []comment.Comment
	Variable node.Node
}

func NewPostInc(Variable node.Node) *PostInc {
	return &PostInc{
		nil,
		nil,
		Variable,
	}
}

func (n *PostInc) Attributes() map[string]interface{} {
	return nil
}

func (n *PostInc) Position() *node.Position {
	return n.position
}

func (n *PostInc) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *PostInc) Comments() []comment.Comment {
	return n.comments
}

func (n *PostInc) SetComments(c []comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *PostInc) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	v.LeaveNode(n)
}
