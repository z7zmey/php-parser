package expr

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Exit node
type Exit struct {
	Meta     []meta.Meta
	Position *position.Position
	Expr     node.Node
}

// NewExit node constructor
func NewExit(Expr node.Node) *Exit {
	return &Exit{
		Expr: Expr,
	}
}

// SetPosition sets node position
func (n *Exit) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Exit) GetPosition() *position.Position {
	return n.Position
}

func (n *Exit) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *Exit) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *Exit) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Exit) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		v.EnterChildNode("Expr", n)
		n.Expr.Walk(v)
		v.LeaveChildNode("Expr", n)
	}

	v.LeaveNode(n)
}
