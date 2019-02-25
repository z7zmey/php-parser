package expr

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Isset node
type Isset struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Variables    []node.Node
}

// NewIsset node constructor
func NewIsset(Variables []node.Node) *Isset {
	return &Isset{
		FreeFloating: nil,
		Variables:    Variables,
	}
}

// SetPosition sets node position
func (n *Isset) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Isset) GetPosition() *position.Position {
	return n.Position
}

func (n *Isset) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *Isset) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Isset) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variables != nil {
		v.EnterChildList("Variables", n)
		for _, nn := range n.Variables {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Variables", n)
	}

	v.LeaveNode(n)
}
