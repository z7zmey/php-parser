package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// AltFor node
type AltFor struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Init         []node.Node
	Cond         []node.Node
	Loop         []node.Node
	Stmt         node.Node
}

// NewAltFor node constructor
func NewAltFor(Init []node.Node, Cond []node.Node, Loop []node.Node, Stmt node.Node) *AltFor {
	return &AltFor{
		FreeFloating: nil,
		Init:         Init,
		Cond:         Cond,
		Loop:         Loop,
		Stmt:         Stmt,
	}
}

// SetPosition sets node position
func (n *AltFor) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *AltFor) GetPosition() *position.Position {
	return n.Position
}

func (n *AltFor) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *AltFor) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *AltFor) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Init != nil {
		v.EnterChildList("Init", n)
		for _, nn := range n.Init {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Init", n)
	}

	if n.Cond != nil {
		v.EnterChildList("Cond", n)
		for _, nn := range n.Cond {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Cond", n)
	}

	if n.Loop != nil {
		v.EnterChildList("Loop", n)
		for _, nn := range n.Loop {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Loop", n)
	}

	if n.Stmt != nil {
		v.EnterChildNode("Stmt", n)
		n.Stmt.Walk(v)
		v.LeaveChildNode("Stmt", n)
	}

	v.LeaveNode(n)
}
