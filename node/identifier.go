package node

// Identifier node
type Identifier struct {
	Value string
}

// NewIdentifier node constuctor
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
func (n *Identifier) Walk(v Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
