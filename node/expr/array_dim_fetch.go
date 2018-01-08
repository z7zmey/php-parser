package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type ArrayDimFetch struct {
	position *node.Position
	comments *[]comment.Comment
	Variable node.Node
	Dim      node.Node
}

func NewArrayDimFetch(Variable node.Node, Dim node.Node) *ArrayDimFetch {
	return &ArrayDimFetch{
		nil,
		nil,
		Variable,
		Dim,
	}
}

func (n *ArrayDimFetch) Attributes() map[string]interface{} {
	return nil
}

func (n *ArrayDimFetch) Position() *node.Position {
	return n.position
}

func (n *ArrayDimFetch) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *ArrayDimFetch) Comments() *[]comment.Comment {
	return n.comments
}

func (n *ArrayDimFetch) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *ArrayDimFetch) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	if n.Dim != nil {
		vv := v.GetChildrenVisitor("Dim")
		n.Dim.Walk(vv)
	}

	v.LeaveNode(n)
}
