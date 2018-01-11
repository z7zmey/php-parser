package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

// TraitUseAlias node
type TraitUseAlias struct {
	Ref      node.Node
	Modifier node.Node
	Alias    node.Node
}

// NewTraitUseAlias node constuctor
func NewTraitUseAlias(Ref node.Node, Modifier node.Node, Alias node.Node) *TraitUseAlias {
	return &TraitUseAlias{
		Ref,
		Modifier,
		Alias,
	}
}

// Attributes returns node attributes as map
func (n *TraitUseAlias) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *TraitUseAlias) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Ref != nil {
		vv := v.GetChildrenVisitor("Ref")
		n.Ref.Walk(vv)
	}

	if n.Modifier != nil {
		vv := v.GetChildrenVisitor("Modifier")
		n.Modifier.Walk(vv)
	}

	if n.Alias != nil {
		vv := v.GetChildrenVisitor("Alias")
		n.Alias.Walk(vv)
	}

	v.LeaveNode(n)
}
