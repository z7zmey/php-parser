package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// TraitUse node
type TraitUse struct {
	FreeFloating        freefloating.Collection
	Position            *position.Position
	Traits              []node.Node
	TraitAdaptationList node.Node
}

// NewTraitUse node constructor
func NewTraitUse(Traits []node.Node, InnerAdaptationList node.Node) *TraitUse {
	return &TraitUse{
		FreeFloating:        nil,
		Traits:              Traits,
		TraitAdaptationList: InnerAdaptationList,
	}
}

// SetPosition sets node position
func (n *TraitUse) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *TraitUse) GetPosition() *position.Position {
	return n.Position
}

func (n *TraitUse) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *TraitUse) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *TraitUse) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Traits != nil {
		v.EnterChildList("Traits", n)
		for _, nn := range n.Traits {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Traits", n)
	}

	if n.TraitAdaptationList != nil {
		v.EnterChildNode("TraitAdaptationList", n)
		n.TraitAdaptationList.Walk(v)
		v.LeaveChildNode("TraitAdaptationList", n)
	}

	v.LeaveNode(n)
}
