package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// TraitMethodRef node
type TraitMethodRef struct {
	Trait  node.Node
	Method node.Node
}

// NewTraitMethodRef node constructor
func NewTraitMethodRef(Trait node.Node, Method node.Node) *TraitMethodRef {
	return &TraitMethodRef{
		Trait,
		Method,
	}
}

// Attributes returns node attributes as map
func (n *TraitMethodRef) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *TraitMethodRef) Walk(v walker.Visitor) {
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
