package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type TraitUseAlias struct {
	attributes map[string]interface{}
	position   *node.Position
	ref        node.Node
	modifier   node.Node
	alias      node.Node
}

func NewTraitUseAlias(ref node.Node, modifier node.Node, alias node.Node) node.Node {
	return TraitUseAlias{
		map[string]interface{}{},
		nil,
		ref,
		modifier,
		alias,
	}
}

func (n TraitUseAlias) Attributes() map[string]interface{} {
	return n.attributes
}

func (n TraitUseAlias) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n TraitUseAlias) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n TraitUseAlias) Position() *node.Position {
	return n.position
}

func (n TraitUseAlias) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n TraitUseAlias) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.ref != nil {
		vv := v.GetChildrenVisitor("ref")
		n.ref.Walk(vv)
	}

	if n.modifier != nil {
		vv := v.GetChildrenVisitor("modifier")
		n.modifier.Walk(vv)
	}

	if n.alias != nil {
		vv := v.GetChildrenVisitor("alias")
		n.alias.Walk(vv)
	}

	v.LeaveNode(n)
}
