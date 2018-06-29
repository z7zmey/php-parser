package expr

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// BitwiseNot node
type BitwiseNot struct {
	Meta     []meta.Meta
	Position *position.Position
	Expr     node.Node
}

// NewBitwiseNot node constructor
func NewBitwiseNot(Expression node.Node) *BitwiseNot {
	return &BitwiseNot{
		Expr: Expression,
	}
}

// SetPosition sets node position
func (n *BitwiseNot) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *BitwiseNot) GetPosition() *position.Position {
	return n.Position
}

func (n *BitwiseNot) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *BitwiseNot) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *BitwiseNot) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *BitwiseNot) Walk(v walker.Visitor) {
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
