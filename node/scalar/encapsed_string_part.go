package scalar

import (
	"github.com/z7zmey/php-parser/node"
)

// EncapsedStringPart node
type EncapsedStringPart struct {
	Value string
}

// NewEncapsedStringPart node constuctor
func NewEncapsedStringPart(Value string) *EncapsedStringPart {
	return &EncapsedStringPart{
		Value,
	}
}

// Attributes returns node attributes as map
func (n *EncapsedStringPart) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *EncapsedStringPart) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
