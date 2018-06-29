package stmt

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// While node
type While struct {
	Meta     []meta.Meta
	Position *position.Position
	Cond     node.Node
	Stmt     node.Node
}

// NewWhile node constructor
func NewWhile(Cond node.Node, Stmt node.Node) *While {
	return &While{
		Cond: Cond,
		Stmt: Stmt,
	}
}

// SetPosition sets node position
func (n *While) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *While) GetPosition() *position.Position {
	return n.Position
}

func (n *While) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *While) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *While) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *While) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Cond != nil {
		v.EnterChildNode("Cond", n)
		n.Cond.Walk(v)
		v.LeaveChildNode("Cond", n)
	}

	if n.Stmt != nil {
		v.EnterChildNode("Stmt", n)
		n.Stmt.Walk(v)
		v.LeaveChildNode("Stmt", n)
	}

	v.LeaveNode(n)
}
