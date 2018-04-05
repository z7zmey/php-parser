package node

import "github.com/z7zmey/php-parser/walker"

// Identifier node
type Identifier struct {
	Value string
}

// NewIdentifier node constructor
func NewIdentifier(Value string) *Identifier {
	return &Identifier{
		Value,
	}
}

// Attributes returns node attributes as map
func (n *Identifier) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Identifier) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
