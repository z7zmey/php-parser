package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// TraitUsePrecedence node
type TraitUsePrecedence struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Ref          node.Node
	Insteadof    []node.Node
}

// NewTraitUsePrecedence node constructor
func NewTraitUsePrecedence(Ref node.Node, Insteadof []node.Node) *TraitUsePrecedence {
	return &TraitUsePrecedence{
		FreeFloating: nil,
		Ref:          Ref,
		Insteadof:    Insteadof,
	}
}

// SetPosition sets node position
func (n *TraitUsePrecedence) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *TraitUsePrecedence) GetPosition() *position.Position {
	return n.Position
}

func (n *TraitUsePrecedence) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
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
		v.EnterChildNode("Ref", n)
		n.Ref.Walk(v)
		v.LeaveChildNode("Ref", n)
	}

	if n.Insteadof != nil {
		v.EnterChildList("Insteadof", n)
		for _, nn := range n.Insteadof {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Insteadof", n)
	}

	v.LeaveNode(n)
}
