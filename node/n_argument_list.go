package node

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// ArgumentList node
type ArgumentList struct {
	Meta      []meta.Meta
	Position  *position.Position
	Arguments []Node
}

// NewArgumentList node constructor
func NewArgumentList(Arguments []Node) *ArgumentList {
	return &ArgumentList{
		Arguments: Arguments,
	}
}

// SetPosition sets node position
func (n *ArgumentList) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *ArgumentList) GetPosition() *position.Position {
	return n.Position
}

func (n *ArgumentList) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *ArgumentList) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *ArgumentList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ArgumentList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Arguments != nil {
		v.EnterChildList("Arguments", n)
		for _, nn := range n.Arguments {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Arguments", n)
	}

	v.LeaveNode(n)
}
