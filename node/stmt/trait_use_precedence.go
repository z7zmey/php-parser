package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type TraitUsePrecedence struct {
	position  *node.Position
	comments  *[]comment.Comment
	Ref       node.Node
	Insteadof node.Node
}

func NewTraitUsePrecedence(Ref node.Node, Insteadof node.Node) node.Node {
	return &TraitUsePrecedence{
		nil,
		nil,
		Ref,
		Insteadof,
	}
}

func (n TraitUsePrecedence) Attributes() map[string]interface{} {
	return nil
}

func (n TraitUsePrecedence) Position() *node.Position {
	return n.position
}

func (n TraitUsePrecedence) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n TraitUsePrecedence) Comments() *[]comment.Comment {
	return n.comments
}

func (n TraitUsePrecedence) SetComments(c []comment.Comment) node.Node {
	n.comments = &c
	return n
}

func (n TraitUsePrecedence) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Ref != nil {
		vv := v.GetChildrenVisitor("Ref")
		n.Ref.Walk(vv)
	}

	if n.Insteadof != nil {
		vv := v.GetChildrenVisitor("Insteadof")
		n.Insteadof.Walk(vv)
	}

	v.LeaveNode(n)
}
