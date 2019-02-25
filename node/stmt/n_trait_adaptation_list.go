package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// TraitAdaptationList node
type TraitAdaptationList struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Adaptations  []node.Node
}

// NewTraitAdaptationList node constructor
func NewTraitAdaptationList(Adaptations []node.Node) *TraitAdaptationList {
	return &TraitAdaptationList{
		FreeFloating: nil,
		Adaptations:  Adaptations,
	}
}

// SetPosition sets node position
func (n *TraitAdaptationList) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *TraitAdaptationList) GetPosition() *position.Position {
	return n.Position
}

func (n *TraitAdaptationList) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *TraitAdaptationList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *TraitAdaptationList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Adaptations != nil {
		v.EnterChildList("Adaptations", n)
		for _, nn := range n.Adaptations {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Adaptations", n)
	}

	v.LeaveNode(n)
}
