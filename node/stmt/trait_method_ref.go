package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type TraitMethodRef struct {
	position *node.Position
	comments *[]comment.Comment
	Trait    node.Node
	Method   node.Node
}

func NewTraitMethodRef(Trait node.Node, Method node.Node) node.Node {
	return &TraitMethodRef{
		nil,
		nil,
		Trait,
		Method,
	}
}

func (n TraitMethodRef) Attributes() map[string]interface{} {
	return nil
}

func (n TraitMethodRef) Position() *node.Position {
	return n.position
}

func (n TraitMethodRef) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n TraitMethodRef) Comments() *[]comment.Comment {
	return n.comments
}

func (n TraitMethodRef) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n TraitMethodRef) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Trait != nil {
		vv := v.GetChildrenVisitor("Trait")
		n.Trait.Walk(vv)
	}

	if n.Method != nil {
		vv := v.GetChildrenVisitor("Method")
		n.Method.Walk(vv)
	}

	v.LeaveNode(n)
}
