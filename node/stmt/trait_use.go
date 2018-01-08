package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type TraitUse struct {
	position    *node.Position
	comments    []comment.Comment
	Traits      []node.Node
	Adaptations []node.Node
}

func NewTraitUse(Traits []node.Node, Adaptations []node.Node) *TraitUse {
	return &TraitUse{
		nil,
		nil,
		Traits,
		Adaptations,
	}
}

func (n *TraitUse) Attributes() map[string]interface{} {
	return nil
}

func (n *TraitUse) Position() *node.Position {
	return n.position
}

func (n *TraitUse) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *TraitUse) Comments() []comment.Comment {
	return n.comments
}

func (n *TraitUse) SetComments(c []comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *TraitUse) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Traits != nil {
		vv := v.GetChildrenVisitor("Traits")
		for _, nn := range n.Traits {
			nn.Walk(vv)
		}
	}

	if n.Adaptations != nil {
		vv := v.GetChildrenVisitor("Adaptations")
		for _, nn := range n.Adaptations {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
