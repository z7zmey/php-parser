package stmt

import "github.com/z7zmey/php-parser/walker"

// Nop node
type Nop struct {
}

// NewNop node constructor
func NewNop() *Nop {
	return &Nop{}
}

// Attributes returns node attributes as map
func (n *Nop) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Nop) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
