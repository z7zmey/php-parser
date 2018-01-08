package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type TraitUseAlias struct {
	position *node.Position
	comments *[]comment.Comment
	Ref      node.Node
	Modifier node.Node
	Alias    node.Node
}

func NewTraitUseAlias(Ref node.Node, Modifier node.Node, Alias node.Node) *TraitUseAlias {
	return &TraitUseAlias{
		nil,
		nil,
		Ref,
		Modifier,
		Alias,
	}
}

func (n TraitUseAlias) Attributes() map[string]interface{} {
	return nil
}

func (n TraitUseAlias) Position() *node.Position {
	return n.position
}

func (n TraitUseAlias) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n TraitUseAlias) Comments() *[]comment.Comment {
	return n.comments
}

func (n TraitUseAlias) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n TraitUseAlias) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Ref != nil {
		vv := v.GetChildrenVisitor("Ref")
		n.Ref.Walk(vv)
	}

	if n.Modifier != nil {
		vv := v.GetChildrenVisitor("Modifier")
		n.Modifier.Walk(vv)
	}

	if n.Alias != nil {
		vv := v.GetChildrenVisitor("Alias")
		n.Alias.Walk(vv)
	}

	v.LeaveNode(n)
}
