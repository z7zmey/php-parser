package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type ClassConstFetch struct {
	position     *node.Position
	comments     *[]comment.Comment
	Class        node.Node
	ConstantName node.Node
}

func NewClassConstFetch(Class node.Node, ConstantName node.Node) *ClassConstFetch {
	return &ClassConstFetch{
		nil,
		nil,
		Class,
		ConstantName,
	}
}

func (n ClassConstFetch) Attributes() map[string]interface{} {
	return nil
}

func (n ClassConstFetch) Position() *node.Position {
	return n.position
}

func (n ClassConstFetch) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n ClassConstFetch) Comments() *[]comment.Comment {
	return n.comments
}

func (n ClassConstFetch) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n ClassConstFetch) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.ConstantName != nil {
		vv := v.GetChildrenVisitor("ConstantName")
		n.ConstantName.Walk(vv)
	}

	if n.Class != nil {
		vv := v.GetChildrenVisitor("Class")
		n.Class.Walk(vv)
	}

	v.LeaveNode(n)
}
