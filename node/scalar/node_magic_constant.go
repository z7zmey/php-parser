package scalar

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// MagicConstant node
type MagicConstant struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Value        string
}

// NewMagicConstant node constructor
func NewMagicConstant(Value string) *MagicConstant {
	return &MagicConstant{
		FreeFloating: nil,
		Value:        Value,
	}
}

// SetPosition sets node position
func (n *MagicConstant) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *MagicConstant) GetPosition() *position.Position {
	return n.Position
}

func (n *MagicConstant) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *MagicConstant) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *MagicConstant) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
