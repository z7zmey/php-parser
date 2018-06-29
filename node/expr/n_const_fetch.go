package expr

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// ConstFetch node
type ConstFetch struct {
	Meta     []meta.Meta
	Position *position.Position
	Constant node.Node
}

// NewConstFetch node constructor
func NewConstFetch(Constant node.Node) *ConstFetch {
	return &ConstFetch{
		Constant: Constant,
	}
}

// SetPosition sets node position
func (n *ConstFetch) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *ConstFetch) GetPosition() *position.Position {
	return n.Position
}

func (n *ConstFetch) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *ConstFetch) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *ConstFetch) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ConstFetch) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Constant != nil {
		v.EnterChildNode("Constant", n)
		n.Constant.Walk(v)
		v.LeaveChildNode("Constant", n)
	}

	v.LeaveNode(n)
}
