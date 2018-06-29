package stmt

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// AltWhile node
type AltWhile struct {
	Meta     []meta.Meta
	Position *position.Position
	Cond     node.Node
	Stmt     node.Node
}

// NewAltWhile node constructor
func NewAltWhile(Cond node.Node, Stmt node.Node) *AltWhile {
	return &AltWhile{
		Cond: Cond,
		Stmt: Stmt,
	}
}

// SetPosition sets node position
func (n *AltWhile) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *AltWhile) GetPosition() *position.Position {
	return n.Position
}

func (n *AltWhile) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *AltWhile) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *AltWhile) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *AltWhile) Walk(v walker.Visitor) {
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
