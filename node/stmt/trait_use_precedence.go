package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type TraitUsePrecedence struct {
	name       string
	attributes map[string]interface{}
	ref        node.Node
	insteadof  node.Node
}

func NewTraitUsePrecedence(ref node.Node, insteadof node.Node) node.Node {
	return TraitUsePrecedence{
		"TraitUsePrecedence",
		map[string]interface{}{},
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
