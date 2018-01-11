package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

// TraitUsePrecedence node
type TraitUsePrecedence struct {
	Ref       node.Node
	Insteadof node.Node
}

// NewTraitUsePrecedence node constuctor
func NewTraitUsePrecedence(Ref node.Node, Insteadof node.Node) *TraitUsePrecedence {
	return &TraitUsePrecedence{
		Ref,
		Insteadof,
	}
}

// Attributes returns node attributes as map
func (n *TraitUsePrecedence) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *TraitUsePrecedence) Walk(v node.Visitor) {
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
