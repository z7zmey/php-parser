package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type TraitMethodRef struct {
	Trait  node.Node
	Method node.Node
}

func NewTraitMethodRef(Trait node.Node, Method node.Node) *TraitMethodRef {
	return &TraitMethodRef{
		Trait,
		Method,
	}
}

func (n *TraitMethodRef) Attributes() map[string]interface{} {
	return nil
}

func (n *TraitMethodRef) Walk(v node.Visitor) {
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
