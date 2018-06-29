package stmt

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Unset node
type Unset struct {
	Meta     []meta.Meta
	Position *position.Position
	Vars     []node.Node
}

// NewUnset node constructor
func NewUnset(Vars []node.Node) *Unset {
	return &Unset{
		Vars: Vars,
	}
}

// SetPosition sets node position
func (n *Unset) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Unset) GetPosition() *position.Position {
	return n.Position
}

func (n *Unset) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *Unset) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *Unset) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Unset) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Vars != nil {
		v.EnterChildList("Vars", n)
		for _, nn := range n.Vars {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Vars", n)
	}

	v.LeaveNode(n)
}
