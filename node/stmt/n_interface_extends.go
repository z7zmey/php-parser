package stmt

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// InterfaceExtends node
type InterfaceExtends struct {
	Meta           meta.Collection
	Position       *position.Position
	InterfaceNames []node.Node
}

// NewInterfaceExtends node constructor
func NewInterfaceExtends(InterfaceNames []node.Node) *InterfaceExtends {
	return &InterfaceExtends{
		InterfaceNames: InterfaceNames,
	}
}

// SetPosition sets node position
func (n *InterfaceExtends) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *InterfaceExtends) GetPosition() *position.Position {
	return n.Position
}

func (n *InterfaceExtends) GetMeta() *meta.Collection {
	return &n.Meta
}

// Attributes returns node attributes as map
func (n *InterfaceExtends) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *InterfaceExtends) Walk(v walker.Visitor) {
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
