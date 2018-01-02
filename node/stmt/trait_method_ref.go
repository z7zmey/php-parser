package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type TraitMethodRef struct {
	name       string
	attributes map[string]interface{}
	position   *node.Position
	trait      node.Node
	method     node.Node
}

func NewTraitMethodRef(trait node.Node, method node.Node) node.Node {
	return TraitMethodRef{
		"TraitMethodRef",
		map[string]interface{}{},
		nil,
		trait,
		method,
	}
}

func (n TraitMethodRef) Name() string {
	return "TraitMethodRef"
}

func (n TraitMethodRef) Attributes() map[string]interface{} {
	return n.attributes
}

func (n TraitMethodRef) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n TraitMethodRef) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
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

	if n.trait != nil {
		vv := v.GetChildrenVisitor("trait")
		n.trait.Walk(vv)
	}

	if n.method != nil {
		vv := v.GetChildrenVisitor("method")
		n.method.Walk(vv)
	}

	v.LeaveNode(n)
}
