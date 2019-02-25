package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// ClassImplements node
type ClassImplements struct {
	FreeFloating   freefloating.Collection
	Position       *position.Position
	InterfaceNames []node.Node
}

// NewClassImplements node constructor
func NewClassImplements(interfaceNames []node.Node) *ClassImplements {
	return &ClassImplements{
		FreeFloating:   nil,
		InterfaceNames: interfaceNames,
	}
}

// SetPosition sets node position
func (n *ClassImplements) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *ClassImplements) GetPosition() *position.Position {
	return n.Position
}

func (n *ClassImplements) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *ClassImplements) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ClassImplements) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.InterfaceNames != nil {
		v.EnterChildList("InterfaceNames", n)
		for _, nn := range n.InterfaceNames {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("InterfaceNames", n)
	}

	v.LeaveNode(n)
}
