package name

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Relative node
type Relative struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Parts        []node.Node
}

// NewRelative node constructor
func NewRelative(Parts []node.Node) *Relative {
	return &Relative{
		FreeFloating: nil,
		Parts:        Parts,
	}
}

// SetPosition sets node position
func (n *Relative) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Relative) GetPosition() *position.Position {
	return n.Position
}

func (n *Relative) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *Relative) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Relative) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Parts != nil {
		v.EnterChildList("Parts", n)
		for _, nn := range n.Parts {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Parts", n)
	}

	v.LeaveNode(n)
}

// GetParts returns the name parts
func (n *Relative) GetParts() []node.Node {
	return n.Parts
}
