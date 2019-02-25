package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// TraitMethodRef node
type TraitMethodRef struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Trait        node.Node
	Method       node.Node
}

// NewTraitMethodRef node constructor
func NewTraitMethodRef(Trait node.Node, Method node.Node) *TraitMethodRef {
	return &TraitMethodRef{
		FreeFloating: nil,
		Trait:        Trait,
		Method:       Method,
	}
}

// SetPosition sets node position
func (n *TraitMethodRef) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *TraitMethodRef) GetPosition() *position.Position {
	return n.Position
}

func (n *TraitMethodRef) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
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
		v.EnterChildNode("Trait", n)
		n.Trait.Walk(v)
		v.LeaveChildNode("Trait", n)
	}

	if n.Method != nil {
		v.EnterChildNode("Method", n)
		n.Method.Walk(v)
		v.LeaveChildNode("Method", n)
	}

	v.LeaveNode(n)
}
