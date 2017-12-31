package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type TraitUsePrecedence struct {
	name       string
	attributes map[string]interface{}
	position *node.Position
	ref        node.Node
	insteadof  node.Node
}

func NewTraitUsePrecedence(ref node.Node, insteadof node.Node) node.Node {
	return TraitUsePrecedence{
		"TraitUsePrecedence",
		map[string]interface{}{},
		nil,
		ref,
		insteadof,
	}
}

func (n TraitUsePrecedence) Name() string {
	return "TraitUsePrecedence"
}

func (n TraitUsePrecedence) Attributes() map[string]interface{} {
	return n.attributes
}

func (n TraitUsePrecedence) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n TraitUsePrecedence) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n TraitUsePrecedence) Position() *node.Position {
	return n.position
}

func (n TraitUsePrecedence) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n TraitUsePrecedence) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.ref != nil {
		vv := v.GetChildrenVisitor("ref")
		n.ref.Walk(vv)
	}

	if n.insteadof != nil {
		vv := v.GetChildrenVisitor("insteadof")
		n.insteadof.Walk(vv)
	}

	v.LeaveNode(n)
}
