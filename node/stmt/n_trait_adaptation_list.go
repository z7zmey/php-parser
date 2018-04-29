package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// TraitAdaptationList node
type TraitAdaptationList struct {
	Adaptations []node.Node
}

// NewTraitAdaptationList node constructor
func NewTraitAdaptationList(Adaptations []node.Node) *TraitAdaptationList {
	return &TraitAdaptationList{
		Adaptations,
	}
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
		vv := v.GetChildrenVisitor("Adaptations")
		for _, nn := range n.Adaptations {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
