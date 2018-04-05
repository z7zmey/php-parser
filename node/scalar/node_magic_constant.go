package scalar

import "github.com/z7zmey/php-parser/walker"

// MagicConstant node
type MagicConstant struct {
	Value string
}

// NewMagicConstant node constructor
func NewMagicConstant(Value string) *MagicConstant {
	return &MagicConstant{
		Value,
	}
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
