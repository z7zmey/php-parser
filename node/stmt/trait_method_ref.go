package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type TraitMethodRef struct {
	attributes map[string]interface{}
	position   *node.Position
	trait      node.Node
	method     node.Node
}

func NewTraitMethodRef(trait node.Node, method node.Node) node.Node {
	return &TraitMethodRef{
		map[string]interface{}{},
		nil,
		trait,
		method,
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
