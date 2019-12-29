package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// PropertyList node
type PropertyList struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Modifiers    []node.Node
	Type         node.Node
	Properties   []node.Node
}

// NewPropertyList node constructor
func NewPropertyList(Modifiers []node.Node, Type node.Node, Properties []node.Node) *PropertyList {
	return &PropertyList{
		FreeFloating: nil,
		Modifiers:    Modifiers,
		Type:         Type,
		Properties:   Properties,
	}
}

// SetPosition sets node position
func (n *PropertyList) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *PropertyList) GetPosition() *position.Position {
	return n.Position
}

func (n *PropertyList) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *PropertyList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *PropertyList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Modifiers != nil {
		v.EnterChildList("Modifiers", n)
		for _, nn := range n.Modifiers {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Modifiers", n)
	}

	if n.Type != nil {
		v.EnterChildNode("Type", n)
		n.Type.Walk(v)
		v.LeaveChildNode("Type", n)
	}

	if n.Properties != nil {
		v.EnterChildList("Properties", n)
		for _, nn := range n.Properties {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Properties", n)
	}

	v.LeaveNode(n)
}
