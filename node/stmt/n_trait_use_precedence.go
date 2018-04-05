package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// TraitUsePrecedence node
type TraitUsePrecedence struct {
	Ref       node.Node
	Insteadof []node.Node
}

// NewTraitUsePrecedence node constructor
func NewTraitUsePrecedence(Ref node.Node, Insteadof []node.Node) *TraitUsePrecedence {
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
func (n *TraitUsePrecedence) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Ref != nil {
		vv := v.GetChildrenVisitor("Ref")
		n.Ref.Walk(vv)
	}

	if n.Insteadof != nil {
		vv := v.GetChildrenVisitor("Insteadof")
		for _, nn := range n.Insteadof {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
