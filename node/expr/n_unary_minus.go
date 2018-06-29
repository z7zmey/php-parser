package expr

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// UnaryMinus node
type UnaryMinus struct {
	Meta     []meta.Meta
	Position *position.Position
	Expr     node.Node
}

// NewUnaryMinus node constructor
func NewUnaryMinus(Expression node.Node) *UnaryMinus {
	return &UnaryMinus{
		Expr: Expression,
	}
}

// SetPosition sets node position
func (n *UnaryMinus) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *UnaryMinus) GetPosition() *position.Position {
	return n.Position
}

func (n *UnaryMinus) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *UnaryMinus) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *UnaryMinus) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *UnaryMinus) Walk(v walker.Visitor) {
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
