package expr

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Yield node
type Yield struct {
	Meta     meta.Collection
	Position *position.Position
	Key      node.Node
	Value    node.Node
}

// NewYield node constructor
func NewYield(Key node.Node, Value node.Node) *Yield {
	return &Yield{
		Key:   Key,
		Value: Value,
	}
}

// SetPosition sets node position
func (n *Yield) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Yield) GetPosition() *position.Position {
	return n.Position
}

func (n *Yield) GetMeta() *meta.Collection {
	return &n.Meta
}

// Attributes returns node attributes as map
func (n *Yield) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Yield) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Key != nil {
		v.EnterChildNode("Key", n)
		n.Key.Walk(v)
		v.LeaveChildNode("Key", n)
	}

	if n.Value != nil {
		v.EnterChildNode("Value", n)
		n.Value.Walk(v)
		v.LeaveChildNode("Value", n)
	}

	v.LeaveNode(n)
}
