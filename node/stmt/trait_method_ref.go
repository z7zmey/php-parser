package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type TraitMethodRef struct {
	attributes map[string]interface{}
	position   *node.Position
	Trait      node.Node
	Method     node.Node
}

func NewTraitMethodRef(Trait node.Node, Method node.Node) node.Node {
	return &TraitMethodRef{
		map[string]interface{}{},
		nil,
		Trait,
		Method,
	}
}

func (n TraitMethodRef) Attributes() map[string]interface{} {
	return n.attributes
}

func (n TraitMethodRef) Position() *node.Position {
	return n.position
}

func (n TraitMethodRef) SetPosition(p *node.Position) node.Node {
	n.position = p
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
